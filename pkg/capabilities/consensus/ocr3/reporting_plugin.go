package ocr3

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"google.golang.org/protobuf/proto"

	ocrcommon "github.com/goplugin/plugin-libocr/commontypes"
	"github.com/goplugin/plugin-libocr/offchainreporting2/types"
	"github.com/goplugin/plugin-libocr/offchainreporting2plus/ocr3types"

	"github.com/goplugin/plugin-common/pkg/capabilities/consensus/ocr3/requests"

	"google.golang.org/protobuf/types/known/structpb"

	pbtypes "github.com/goplugin/plugin-common/pkg/capabilities/consensus/ocr3/types"
	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/values"
)

var _ ocr3types.ReportingPlugin[[]byte] = (*reportingPlugin)(nil)

type capabilityIface interface {
	getAggregator(workflowID string) (pbtypes.Aggregator, error)
	getEncoder(workflowID string) (pbtypes.Encoder, error)
	getRegisteredWorkflowsIDs() []string
	unregisterWorkflowID(workflowID string)
}

// TODO: 3,600 is the amount of rounds we allow as threshold. This should be configurable.
// This is affected by OCR round time.
const outcomePruningThreshold = 3600

type reportingPlugin struct {
	batchSize int
	s         *requests.Store
	r         capabilityIface
	config    ocr3types.ReportingPluginConfig
	lggr      logger.Logger
}

func newReportingPlugin(s *requests.Store, r capabilityIface, batchSize int, config ocr3types.ReportingPluginConfig, lggr logger.Logger) (*reportingPlugin, error) {
	// TODO: extract limits from OnchainConfig
	// and perform validation.

	return &reportingPlugin{
		s:         s,
		r:         r,
		batchSize: batchSize,
		config:    config,
		lggr:      logger.Named(lggr, "OCR3ConsensusReportingPlugin"),
	}, nil
}

func (r *reportingPlugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	batch, err := r.s.FirstN(r.batchSize)
	if err != nil {
		r.lggr.Errorw("could not retrieve batch", "error", err)
		return nil, err
	}

	ids := []*pbtypes.Id{}
	allExecutionIDs := []string{}
	for _, rq := range batch {
		ids = append(ids, &pbtypes.Id{
			WorkflowExecutionId:      rq.WorkflowExecutionID,
			WorkflowId:               rq.WorkflowID,
			WorkflowOwner:            rq.WorkflowOwner,
			WorkflowName:             rq.WorkflowName,
			WorkflowDonId:            rq.WorkflowDonID,
			WorkflowDonConfigVersion: rq.WorkflowDonConfigVersion,
			ReportId:                 rq.ReportID,
			KeyId:                    rq.KeyID,
		})
		allExecutionIDs = append(allExecutionIDs, rq.WorkflowExecutionID)
	}

	r.lggr.Debugw("Query complete", "len", len(ids), "allExecutionIDs", allExecutionIDs)
	return proto.MarshalOptions{Deterministic: true}.Marshal(&pbtypes.Query{
		Ids: ids,
	})
}

func (r *reportingPlugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	queryReq := &pbtypes.Query{}
	err := proto.Unmarshal(query, queryReq)
	if err != nil {
		return nil, err
	}

	weids := []string{}
	for _, q := range queryReq.Ids {
		if q == nil {
			r.lggr.Debugw("skipping nil id for query", "query", queryReq)
			continue
		}
		weids = append(weids, q.WorkflowExecutionId)
	}

	reqs := r.s.GetByIDs(weids)
	reqMap := map[string]*requests.Request{}
	for _, req := range reqs {
		reqMap[req.WorkflowExecutionID] = req
	}

	obs := &pbtypes.Observations{}
	allExecutionIDs := []string{}
	for _, weid := range weids {
		rq, ok := reqMap[weid]
		if !ok {
			r.lggr.Debugw("could not find local observations for weid requested in the query", "weid", weid)
			continue
		}
		listProto := values.Proto(rq.Observations).GetListValue()
		if listProto == nil {
			r.lggr.Errorw("observations are not a list", "weID", rq.WorkflowExecutionID)
			continue
		}
		newOb := &pbtypes.Observation{
			Observations: listProto,
			Id: &pbtypes.Id{
				WorkflowExecutionId:      rq.WorkflowExecutionID,
				WorkflowId:               rq.WorkflowID,
				WorkflowOwner:            rq.WorkflowOwner,
				WorkflowName:             rq.WorkflowName,
				WorkflowDonId:            rq.WorkflowDonID,
				WorkflowDonConfigVersion: rq.WorkflowDonConfigVersion,
				ReportId:                 rq.ReportID,
				KeyId:                    rq.KeyID,
			},
		}

		obs.Observations = append(obs.Observations, newOb)
		allExecutionIDs = append(allExecutionIDs, rq.WorkflowExecutionID)
	}
	obs.RegisteredWorkflowIds = r.r.getRegisteredWorkflowsIDs()

	r.lggr.Debugw("Observation complete", "len", len(obs.Observations), "queryLen", len(queryReq.Ids), "allExecutionIDs", allExecutionIDs)
	return proto.MarshalOptions{Deterministic: true}.Marshal(obs)
}

func (r *reportingPlugin) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	return nil
}

func (r *reportingPlugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (r *reportingPlugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	// execution ID -> oracle ID -> list of observations
	m := map[string]map[ocrcommon.OracleID][]values.Value{}
	seenWorkflowIDs := map[string]int{}
	for _, o := range aos {
		obs := &pbtypes.Observations{}
		err := proto.Unmarshal(o.Observation, obs)
		if err != nil {
			r.lggr.Errorw("could not unmarshal observation", "error", err, "observation", obs)
			continue
		}

		countedWorkflowIDs := map[string]bool{}
		for _, id := range obs.RegisteredWorkflowIds {
			// Skip if we've already counted this workflow ID. we want to avoid duplicates in the seen workflow IDs.
			if _, ok := countedWorkflowIDs[id]; ok {
				continue
			}

			// Count how many times a workflow ID is seen from Observations, no need for initial value since it's 0 by default.
			seenWorkflowIDs[id]++

			countedWorkflowIDs[id] = true
		}

		for _, rq := range obs.Observations {
			if rq == nil {
				r.lggr.Debugw("skipping nil request in observations", "observations", obs.Observations)
				continue
			}

			if rq.Id == nil {
				r.lggr.Debugw("skipping nil id in request", "request", rq)
				continue
			}

			weid := rq.Id.WorkflowExecutionId

			obsList, innerErr := values.FromListValueProto(rq.Observations)
			if obsList == nil || innerErr != nil {
				r.lggr.Errorw("observations are not a list", "weID", weid, "oracleID", o.Observer, "err", innerErr)
				continue
			}

			if _, ok := m[weid]; !ok {
				m[weid] = make(map[ocrcommon.OracleID][]values.Value)
			}
			m[weid][o.Observer] = obsList.Underlying
		}
	}

	q := &pbtypes.Query{}
	err := proto.Unmarshal(query, q)
	if err != nil {
		return nil, err
	}

	o := &pbtypes.Outcome{}
	err = proto.Unmarshal(outctx.PreviousOutcome, o)
	if err != nil {
		return nil, err
	}
	if o.Outcomes == nil {
		o.Outcomes = map[string]*pbtypes.AggregationOutcome{}
	}

	// Wipe out the CurrentReports. This gets regenerated
	// every time since we only want to transmit reports that
	// are part of the current Query.
	o.CurrentReports = []*pbtypes.Report{}
	var allExecutionIDs []string

	for _, weid := range q.Ids {
		if weid == nil {
			r.lggr.Debugw("skipping nil id in query", "query", q)
			continue
		}
		lggr := logger.With(r.lggr, "executionID", weid.WorkflowExecutionId, "workflowID", weid.WorkflowId)
		obs, ok := m[weid.WorkflowExecutionId]
		if !ok {
			r.lggr.Debugw("could not find any observations matching weid requested in the query", "weid", weid.WorkflowExecutionId)
			continue
		}

		workflowOutcome, ok := o.Outcomes[weid.WorkflowId]
		if !ok {
			r.lggr.Debugw("could not find existing outcome for workflow, aggregator will create a new one", "workflowID", weid.WorkflowId)
		}

		if len(obs) < (2*r.config.F + 1) {
			r.lggr.Debugw("insufficient observations for workflow execution id", "weid", weid.WorkflowExecutionId)
			continue
		}

		agg, err2 := r.r.getAggregator(weid.WorkflowId)
		if err2 != nil {
			r.lggr.Errorw("could not retrieve aggregator for workflow", "error", err, "workflowID", weid.WorkflowId)
			continue
		}

		outcome, err2 := agg.Aggregate(workflowOutcome, obs, r.config.F)
		if err2 != nil {
			r.lggr.Errorw("error aggregating outcome", "error", err, "workflowID", weid.WorkflowId)
			return nil, err
		}

		// Only if the previous outcome exists:
		// We carry the last seen round from the previous outcome, since the aggregation does carry it.
		// So each `Aggregate()` call will return an outcome with a zero value for LastSeenAt.
		if workflowOutcome != nil {
			outcome.LastSeenAt = workflowOutcome.LastSeenAt
		}

		report := &pbtypes.Report{
			Outcome: outcome,
			Id:      weid,
		}
		o.CurrentReports = append(o.CurrentReports, report)
		allExecutionIDs = append(allExecutionIDs, weid.WorkflowExecutionId)

		o.Outcomes[weid.WorkflowId] = outcome
	}

	// We need to prune outcomes from previous workflows that are no longer relevant.
	for workflowID, outcome := range o.Outcomes {
		// Update the last seen round for this outcome. But this should only happen if the workflow is seen by F+1 nodes.
		if seenWorkflowIDs[workflowID] >= (r.config.F + 1) {
			r.lggr.Debugw("updating last seen round of outcome for workflow", "workflowID", workflowID)
			outcome.LastSeenAt = outctx.SeqNr
		} else if outctx.SeqNr-outcome.LastSeenAt > outcomePruningThreshold {
			r.lggr.Debugw("pruning outcome for workflow", "workflowID", workflowID, "SeqNr", outctx.SeqNr, "lastSeenAt", outcome.LastSeenAt)
			delete(o.Outcomes, workflowID)
			r.r.unregisterWorkflowID(workflowID)
		}
	}

	rawOutcome, err := proto.MarshalOptions{Deterministic: true}.Marshal(o)
	h := sha256.New()
	h.Write(rawOutcome)
	outcomeHash := h.Sum(nil)
	r.lggr.Debugw("Outcome complete", "len", len(o.Outcomes), "nAggregatedWorkflowExecutions", len(o.CurrentReports), "allExecutionIDs", allExecutionIDs, "outcomeHash", hex.EncodeToString(outcomeHash), "err", err)
	return rawOutcome, err
}

func marshalReportInfo(info *pbtypes.ReportInfo, keyID string) ([]byte, error) {
	p, err := proto.MarshalOptions{Deterministic: true}.Marshal(info)
	if err != nil {
		return nil, err
	}

	infos, err := structpb.NewStruct(map[string]any{
		"keyBundleName": keyID,
		"reportInfo":    p,
	})
	if err != nil {
		return nil, err
	}

	ip, err := proto.MarshalOptions{Deterministic: true}.Marshal(infos)
	if err != nil {
		return nil, err
	}

	return ip, nil
}

func (r *reportingPlugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	o := &pbtypes.Outcome{}
	err := proto.Unmarshal(outcome, o)
	if err != nil {
		return nil, err
	}

	reports := []ocr3types.ReportWithInfo[[]byte]{}

	for _, report := range o.CurrentReports {
		r.lggr.Debugw("generating reports", "len", len(o.CurrentReports), "shouldReport", report.Outcome.ShouldReport, "executionID", report.Id.WorkflowExecutionId)

		lggr := logger.With(
			r.lggr,
			"workflowID", report.Id.WorkflowId,
			"executionID", report.Id.WorkflowExecutionId,
			"shouldReport", report.Outcome.ShouldReport,
		)


		outcome, id := report.Outcome, report.Id

		info := &pbtypes.ReportInfo{
			Id:           id,
			ShouldReport: outcome.ShouldReport,
		}

		var report []byte
		if info.ShouldReport {
			meta := &pbtypes.Metadata{
				Version:          1,
				ExecutionID:      id.WorkflowExecutionId,
				Timestamp:        0, // TODO include timestamp in consensus phase
				DONID:            id.WorkflowDonId,
				DONConfigVersion: id.WorkflowDonConfigVersion,
				WorkflowID:       id.WorkflowId,
				WorkflowName:     id.WorkflowName,
				WorkflowOwner:    id.WorkflowOwner,
				ReportID:         id.ReportId,
			}
			newOutcome, err := pbtypes.AppendMetadata(outcome, meta)
			if err != nil {
				r.lggr.Errorw("could not append IDs")
				continue
			}

			enc, err := r.r.getEncoder(id.WorkflowId)
			if err != nil {
				r.lggr.Errorw("could not retrieve encoder for workflow", "error", err, "workflowID", id.WorkflowId)
				continue
			}

			mv, err := values.FromMapValueProto(newOutcome.EncodableOutcome)
			if err != nil {
				r.lggr.Errorw("could not decode map from map value proto", "error", err)
				continue
			}

			report, err = enc.Encode(context.TODO(), *mv)
			if err != nil {
				r.lggr.Errorw("could not encode report for workflow", "error", err, "workflowID", id.WorkflowId)
				continue
			}
		}

		p, err := proto.MarshalOptions{Deterministic: true}.Marshal(info)
		if err != nil {
			r.lggr.Errorw("could not marshal id into ReportWithInfo", "error", err, "workflowID", id.WorkflowId, "shouldReport", info.ShouldReport)
			continue
		}

		// Append every report, even if shouldReport = false, to let the transmitter mark the step as complete.
		reports = append(reports, ocr3types.ReportWithInfo[[]byte]{
			Report: report,
			Info:   infob,
		})
	}

	r.lggr.Debugw("Reports complete", "len", len(reports))
	return reports, nil
}

func (r *reportingPlugin) ShouldAcceptAttestedReport(ctx context.Context, seqNr uint64, rwi ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	// True because we always want to transmit a report, even if shouldReport = false.
	return true, nil
}

func (r *reportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, seqNr uint64, rwi ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	// True because we always want to transmit a report, even if shouldReport = false.
	return true, nil
}

func (r *reportingPlugin) Close() error {
	return nil
}
