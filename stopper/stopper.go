package stopper

import (
	"context"
	"sync"
)

type Stopper struct {
	sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewStopper() *Stopper {
	ctx, cancel := context.WithCancel(context.Background())
	return &Stopper{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Stopper) Run(f func(context.Context)) {
	s.Add(1)
	go func(ctx context.Context) {
		f(ctx)
		s.Done()
	}(s.ctx)
}

func (s *Stopper) Stop() {
	s.cancel()
	s.Wait()
}
