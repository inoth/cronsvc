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

	receiver chan []byte
	execute  chan []byte
}

func New(opts ...Option) *Executor {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}
	ecr = &Executor{
		option:   o,
		receiver: make(chan []byte, 1000),
		execute:  make(chan []byte, 1000),
	}
	return ecr
}

func (e *Executor) Name() string {
	return name
}

func (e *Executor) Start(ctx context.Context) error {

	sctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go e.pipline(sctx)

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-e.execute:
			fmt.Printf("run msg %s\n", string(msg))
		}
	}
	return nil
}

func (e *Executor) Stop(ctx context.Context) error {
	return nil
}

func (e *Executor) pipline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-e.receiver:
			//TODO: do something
			e.execute <- msg
		}
	}
}
