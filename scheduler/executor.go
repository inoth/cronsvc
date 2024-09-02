package scheduler

import (
	"context"
	"sync"
)

// https://github.com/holdno/gopherCron/blob/master/agent/scheduler.go

type Scheduler struct {
	option

	m sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	tasks map[string]TaskEvent
}

func New(opts ...Option) *Scheduler {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}
	return &Scheduler{
		option: o,
	}
}
