package executor

import (
	"context"
	"sync"

	"github.com/inoth/cronsvc/internal/util"
)

type Executor struct {
	reloadCount  int
	curTaskCount int
	mu           sync.Mutex

	ctx    context.Context
	cancel func()

	opt option
}

func NewExecutor(opts ...Option) *Executor {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}
	o.id = util.UUID()
	if o.name == "" {
		o.name = o.id
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Executor{
		opt:          o,
		ctx:          ctx,
		cancel:       cancel,
		reloadCount:  0,
		curTaskCount: 0,
	}
}

func (e *Executor) AddTaskCount() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.curTaskCount += 1
}

func (e *Executor) Run() error {
	return nil
}

func (e *Executor) Stop() {
	e.cancel()
}

func (e *Executor) Reload() error {
	e.reloadCount += 1
	return nil
}
