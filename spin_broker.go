package spin

import (
	"context"
	"log"
	"net/http"

	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/db"
)

// spin-broker service example implementation.
// The example methods log the requests and return zero values.
type spinBrokersrvc struct {
	logger *log.Logger
	db     *db.DB
}

// NewSpinBroker returns the spin-broker service implementation.
func NewSpinBroker(logger *log.Logger, dbName string) (spinbroker.Service, error) {
	db, err := db.New(dbName)
	if err != nil {
		return nil, err
	}
	return &spinBrokersrvc{logger: logger, db: db}, nil
}

// New implements new.
func (s *spinBrokersrvc) New(ctx context.Context) (res string, err error) {
	return s.db.NewPackage()
}

// Add
func (s *spinBrokersrvc) Add(ctx context.Context, p *spinbroker.AddPayload) (string, error) {
	return s.db.AddToPackage(p.ID, &db.AddCommand{
		Resource:   p.Resource,
		Action:     p.Action,
		Parameters: p.Parameters,
	})
}

// Enqueue
func (s *spinBrokersrvc) Enqueue(ctx context.Context, p *spinbroker.EnqueuePayload) ([]string, error) {
	return s.db.EnqueuePackage(p.ID)
}

// Enqueued
func (s *spinBrokersrvc) Enqueued(ctx context.Context, p *spinbroker.EnqueuedPayload) (bool, error) {
	return false, http.ErrNotSupported
}

// Status
func (s *spinBrokersrvc) Status(ctx context.Context, p *spinbroker.StatusPayload) (*spinbroker.StatusResult, error) {
	return nil, http.ErrNotSupported
}

// Next
func (s *spinBrokersrvc) Next(ctx context.Context, p *spinbroker.NextPayload) (*spinbroker.NextResult, error) {
	return nil, http.ErrNotSupported
}

// Complete
func (s *spinBrokersrvc) Complete(ctx context.Context, p *spinbroker.CompletePayload) error {
	return http.ErrNotSupported
}
