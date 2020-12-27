package spin

import (
	"context"
	"log"

	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/broker"
)

type spinBrokersrvc struct {
	logger *log.Logger
	db     *broker.DB
}

// NewSpinBroker returns the spin-broker service implementation.
func NewSpinBroker(logger *log.Logger, dbpath string) (spinbroker.Service, error) {
	db, err := broker.New(dbpath)
	if err != nil {
		return nil, err
	}
	return &spinBrokersrvc{logger: logger, db: db}, nil
}

// New
func (s *spinBrokersrvc) New(ctx context.Context) (res string, err error) {
	pkg, err := s.db.NewPackage()
	if err != nil {
		return "", err
	}

	return pkg.UUID(), nil
}

// Add
func (s *spinBrokersrvc) Add(ctx context.Context, p *spinbroker.AddPayload) (string, error) {
	pkg, err := s.db.Package(p.ID)
	if err != nil {
		return "", err
	}

	cmd := &broker.Command{
		Resource:   p.Resource,
		Action:     p.Action,
		Parameters: p.Parameters,
	}
	if err := pkg.Add(cmd); err != nil {
		return "", err
	}

	return cmd.UUID, nil
}

// Enqueue
func (s *spinBrokersrvc) Enqueue(ctx context.Context, p *spinbroker.EnqueuePayload) ([]string, error) {
	pkg, err := s.db.Package(p.ID)
	if err != nil {
		return nil, err
	}

	uuids := []string{}
	list, err := pkg.List()
	if err != nil {
		return nil, err
	}

	for _, c := range list {
		uuids = append(uuids, c.UUID)
	}

	return uuids, pkg.Enqueue()
}

// Status
func (s *spinBrokersrvc) Status(ctx context.Context, p *spinbroker.StatusPayload) (*spinbroker.StatusResult, error) {
	pkg, err := s.db.Package(p.ID)
	if err != nil {
		return nil, err
	}

	if err := pkg.Finished(); err != nil {
		if e, ok := err.(broker.ErrorStatus); ok {
			return &spinbroker.StatusResult{Reason: &e.Reason}, nil
		}
		return nil, err
	}

	return &spinbroker.StatusResult{Status: true}, nil
}

// Next
func (s *spinBrokersrvc) Next(ctx context.Context, p *spinbroker.NextPayload) (*spinbroker.NextResult, error) {
	queue, err := s.db.Queue(p.Resource)
	if err != nil {
		return nil, err
	}

	c, err := queue.Next()
	if err != nil {
		return nil, err
	}

	return &spinbroker.NextResult{
		UUID:       c.UUID,
		Resource:   c.Resource,
		Action:     c.Action,
		Parameters: c.Parameters,
	}, nil
}

// Complete
func (s *spinBrokersrvc) Complete(ctx context.Context, p *spinbroker.CompletePayload) error {
	sr := ""
	if p.StatusReason != nil {
		sr = *p.StatusReason
	}
	return s.db.FinishCommand(p.ID, p.Status, sr)
}
