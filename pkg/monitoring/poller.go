package monitoring

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/goplugin/plugin-common/pkg/services"
	"github.com/goplugin/plugin-common/pkg/timeutil"
)

// Poller implements Updater by periodically invoking a Source's Fetch() method.
type Poller interface {
	Updater // Poller is just another name for updater.
}

// NewSourcePoller builds Pollers for Sources.
// If the Source's Fetch() returns an error it will be reported.
// If it panics, the panic will be recovered and reported as an error and the program will resume operation.
// If the error is ErrNoUpdate, it will not be reported and the Poller will skip this round.
func NewSourcePoller(
	source Source,
	log Logger,
	pollInterval time.Duration,
	fetchTimeout time.Duration,
	bufferCapacity uint32,
) Poller {
	return &sourcePoller{
		log,
		source,
		make(chan interface{}, bufferCapacity),
		pollInterval,
		fetchTimeout,
	}
}

type sourcePoller struct {
	log     Logger
	source  Source
	updates chan interface{}

	pollInterval time.Duration
	fetchTimeout time.Duration
}

// Run should be executed as a goroutine
func (s *sourcePoller) Run(ctx context.Context) {
	s.log.Debugw("poller started")
	defer s.log.Debugw("poller closed")
	// Initial fetch.
	data, err := s.executeFetch(ctx)
	if err != nil {
		if errors.Is(err, ErrNoUpdate) {
			s.log.Debugw("no update found on initial fetch")
		} else if errors.Is(err, context.Canceled) {
			return
		} else {
			s.log.Errorw("failed initial fetch", "error", err)
		}
	} else {
		select {
		case s.updates <- data:
		case <-ctx.Done():
			return
		}
	}

	ticker := services.TickerConfig{
		Initial:   timeutil.JitterPct(0.1).Apply(s.pollInterval),
		JitterPct: 0.2,
	}.NewTicker(s.pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			data, err := s.executeFetch(ctx)
			if err != nil {
				if errors.Is(err, ErrNoUpdate) {
					s.log.Debugw("no update found")
					ticker.Reset()
					continue
				} else if errors.Is(err, context.Canceled) {
					return
				}
				s.log.Errorw("failed to fetch from source", "error", err)
				ticker.Reset()
				continue
			}
			select {
			case s.updates <- data:
			case <-ctx.Done():
				return
			}
			ticker.Reset()
		case <-ctx.Done():
			return
		}
	}
}

func (s *sourcePoller) Updates() <-chan interface{} {
	return s.updates
}

// executeFetch runs Source#Fetch() with a timeout.
// It also captures the error if Fetch() panics and returns it.
func (s *sourcePoller) executeFetch(ctx context.Context) (data interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.fetchTimeout)
	defer cancel()
	defer func() {
		if recoveredErr := recover(); recoveredErr != nil {
			err = fmt.Errorf("Fetch() panicked: %v", recoveredErr)
		}
	}()
	data, err = s.source.Fetch(ctx)
	return data, err
}
