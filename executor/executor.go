package executor

import (
	"context"
	"fmt"
	"sync"
)

const (
	name = "executor"
)

var (
	ecr *Executor
)

type Executor struct {
	option

	mu sync.Mutex

	taskCount int

	receiver chan TaskBody
	execute  chan TaskBody

	tasks map[string]TaskBody
}

func New(opts ...Option) *Executor {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}
	ecr = &Executor{
		option:   o,
		receiver: make(chan TaskBody, 1000),
		execute:  make(chan TaskBody, 1000),
	}
	return ecr
}

func (e *Executor) Name() string {
	return name
}

func (e *Executor) CurrentTaskCount() int {
	return e.taskCount
}

func (e *Executor) TaskChange(add int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.taskCount += add
}

func (e *Executor) AddTask(task TaskBody) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.tasks[task.ID]; !ok {
		e.tasks[task.ID] = task
	}
}

func (e *Executor) Remove(task TaskBody) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.tasks, task.ID)
}

func (e *Executor) Start(ctx context.Context) error {

	sctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go e.pipline(sctx)

	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil && err != context.Canceled {
				return err
			}
			return nil
		case msg := <-e.execute:
			fmt.Printf("run msg %s\n", msg.String())
		}
	}
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
