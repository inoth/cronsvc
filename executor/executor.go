package executor

import (
	"context"
	"fmt"
)

const (
	name = "executor"
)

var (
	ecr *Executor
)

type Executor struct {
	option

	//TODO: 新建pipline处理接收的任务
	receiver chan<- []byte
	execute  <-chan []byte
}

func New(opts ...Option) *Executor {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}
	ecr = &Executor{
		option:   o,
		receiver: make(chan<- []byte, 1000),
		execute:  make(<-chan []byte, 1000),
	}
	return ecr
}

func (e *Executor) Name() string {
	return name
}

func (e *Executor) Start(ctx context.Context) error {
	fmt.Printf("%s start run, max_task %d\n", name, e.MaxTask)
	return nil
}

func (e *Executor) Stop(ctx context.Context) error {
	return nil
}
