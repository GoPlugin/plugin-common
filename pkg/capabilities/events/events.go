package events

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/goplugin/plugin-common/pkg/beholder"
	"github.com/goplugin/plugin-common/pkg/beholder/pb"
	"github.com/goplugin/plugin-common/pkg/values"
)

const (
	// Duplicates the attributes in beholder/message.go::Metadata
	LabelWorkflowOwner             = "workflow_owner_address"
	LabelWorkflowID                = "workflow_id"
	LabelWorkflowExecutionID       = "workflow_execution_id"
	LabelWorkflowName              = "workflow_name"
	LabelCapabilityContractAddress = "capability_contract_address"
	LabelCapabilityID              = "capability_id"
	LabelCapabilityVersion         = "capability_version"
	LabelCapabilityName            = "capability_name"
)

type EmitMetadata struct {
	WorkflowOwner string // required
	WorkflowID    string // required
	WorkflowName  string // required

	WorkflowExecutionID       string // optional
	CapabilityContractAddress string // optional
	CapabilityID              string // optional
	CapabilityVersion         string // optional
	CapabilityName            string // optional
}

func (e EmitMetadata) merge(otherE EmitMetadata) EmitMetadata {
	owner := e.WorkflowOwner
	if otherE.WorkflowOwner != "" {
		owner = otherE.WorkflowOwner
	}

	wid := e.WorkflowID
	if otherE.WorkflowID != "" {
		wid = otherE.WorkflowID
	}

	eid := e.WorkflowExecutionID
	if otherE.WorkflowExecutionID != "" {
		eid = otherE.WorkflowExecutionID
	}

	name := e.WorkflowName
	if otherE.WorkflowName != "" {
		name = otherE.WorkflowName
	}

	addr := e.CapabilityContractAddress
	if otherE.CapabilityContractAddress != "" {
		addr = otherE.CapabilityContractAddress
	}

	capID := e.CapabilityID
	if otherE.CapabilityID != "" {
		capID = otherE.CapabilityID
	}

	capVersion := e.CapabilityVersion
	if otherE.CapabilityVersion != "" {
		capVersion = otherE.CapabilityVersion
	}

	capName := e.CapabilityName
	if otherE.CapabilityName != "" {
		capName = otherE.CapabilityName
	}

	return EmitMetadata{
		WorkflowOwner:             owner,
		WorkflowID:                wid,
		WorkflowExecutionID:       eid,
		WorkflowName:              name,
		CapabilityContractAddress: addr,
		CapabilityID:              capID,
		CapabilityVersion:         capVersion,
		CapabilityName:            capName,
	}
}

func (e EmitMetadata) attrs() []any {
	a := []any{}

	if e.WorkflowOwner != "" {
		a = append(a, LabelWorkflowOwner, e.WorkflowOwner)
	}

	if e.WorkflowID != "" {
		a = append(a, LabelWorkflowID, e.WorkflowID)
	}

	if e.WorkflowExecutionID != "" {
		a = append(a, LabelWorkflowExecutionID, e.WorkflowExecutionID)
	}

	if e.WorkflowName != "" {
		a = append(a, LabelWorkflowName, e.WorkflowName)
	}

	if e.CapabilityContractAddress != "" {
		a = append(a, LabelCapabilityContractAddress, e.CapabilityContractAddress)
	}

	if e.CapabilityID != "" {
		a = append(a, LabelCapabilityID, e.CapabilityID)
	}

	if e.CapabilityVersion != "" {
		a = append(a, LabelCapabilityVersion, e.CapabilityVersion)
	}

	if e.CapabilityName != "" {
		a = append(a, LabelCapabilityName, e.CapabilityName)
	}

	return a
}

type Emitter struct {
	client beholder.Emitter
	md     EmitMetadata
}

func (e *Emitter) With(md EmitMetadata) *Emitter {
	nmd := e.md.merge(md)
	return &Emitter{
		client: e.client,
		md:     nmd,
	}
}

func NewEmitter() *Emitter {
	return &Emitter{
		client: beholder.GetClient().Emitter,
	}
}

type Message struct {
	Msg      string
	Labels   map[string]any
	Metadata EmitMetadata
}

func (e *Emitter) Emit(ctx context.Context, msg Message) error {
	nmd := e.md.merge(msg.Metadata)

	if nmd.WorkflowOwner == "" {
		return errors.New("must provide workflow owner to emit event")
	}

	if nmd.WorkflowID == "" {
		return errors.New("must provide workflow id to emit event")
	}

	if nmd.WorkflowName == "" {
		return errors.New("must provide workflow name to emit event")
	}

	wm, err := values.WrapMap(msg.Labels)
	if err != nil {
		return fmt.Errorf("could not wrap map: %w", err)
	}

	pm := values.ProtoMap(wm)

	bytes, err := proto.Marshal(&pb.BaseMessage{
		Labels: pm,
		Msg:    msg.Msg,
	})
	if err != nil {
		return fmt.Errorf("could not marshal operational event: %w", err)
	}

	attrs := []any{
		"beholder_data_schema",
		"/capabilities-operational-event/versions/1",
		"beholder_data_type",
		"custom_message",
	}

	attrs = append(attrs, nmd.attrs()...)

	return e.client.Emit(
		ctx,
		bytes,
		attrs...,
	)
}
