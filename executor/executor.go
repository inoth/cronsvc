package executor

import (
	"context"
	"fmt"
	"sync"

	"github.com/inoth/cronsvc/executor/collector"
	"github.com/inoth/cronsvc/executor/task"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

const (
	name = "executor"
)

var (
	ecr *Executor
)

type Executor struct {
	option

	c *cron.Cron

	mu     sync.Mutex
	ctx    context.Context
	cancel func()

	col      chan collector.Collector
	receiver chan TaskBody
	execute  chan TaskBody

	tasks map[string]TaskBody
}

func New(opts ...Option) *Executor {
	o := option{
		CollectorCount: 100,
		ReceiverCount:  100,
		ExecuteCount:   100,
	}
	for _, opt := range opts {
		opt(&o)
	}
	ecr = &Executor{
		option:   o,
		c:        cron.New(cron.WithSeconds()),
		col:      make(chan collector.Collector, o.CollectorCount),
		receiver: make(chan TaskBody, o.ReceiverCount),
		execute:  make(chan TaskBody, o.ExecuteCount),
		tasks:    make(map[string]TaskBody),
	}
	return ecr
}

func (e *Executor) Name() string {
	return name
}

func (e *Executor) CurrentTaskCount() int {
	return len(e.c.Entries())
}

func (e *Executor) AddTask(taskBody TaskBody) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.tasks[taskBody.ID]; !ok {

		val, ok := task.TaskMap[taskBody.Tag]
		if !ok {
			return fmt.Errorf("invalid task tag = %s", taskBody.Tag)
		}

		ctx := context.WithValue(e.ctx, "taskId", taskBody.ID)
		taskBody.ctx, taskBody.cancel = context.WithCancel(ctx)

		runID, err := e.c.AddFunc(taskBody.Crontab, func() {
			val().Run(taskBody.ctx, e.col, taskBody.Body)
		})
		if err != nil {
			return errors.Wrap(err, "append task job err")
		}
		taskBody.runID = int(runID)

		e.tasks[taskBody.ID] = taskBody

		fmt.Printf("start task %s success; current task %d\n", taskBody.ID, e.CurrentTaskCount())
	}
	return nil
}

func (e *Executor) RemoveTask(taskId string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if task, ok := e.tasks[taskId]; ok {
		task.cancel()
		delete(e.tasks, taskId)

		e.c.Remove(cron.EntryID(task.runID))

		fmt.Printf("remove task %s success; current task %d\n", taskId, e.CurrentTaskCount())
	}
}

func (e *Executor) Start(ctx context.Context) error {

	e.ctx, e.cancel = context.WithCancel(ctx)
	defer e.cancel()

	go e.pipline()
	go e.runCollector()
	e.c.Start()

	for {
		select {
		case <-e.ctx.Done():
			if err := e.ctx.Err(); err != nil && err != context.Canceled {
				e.c.Stop()
				return err
			}
			return nil
		case msg := <-e.execute:
			if err := e.AddTask(msg); err != nil {
				fmt.Printf("add task err %v\n", err)
			}
		}
	}
}

func (e *Executor) Stop(ctx context.Context) error {
	e.cancel()
	return nil
}

func (e *Executor) pipline() {
	for {
		select {
		case <-e.ctx.Done():
			return
		case msg := <-e.receiver:
			//TODO: do something
			e.execute <- msg
		}
	}
}

func (e *Executor) runCollector() {
	for {
		select {
		case <-e.ctx.Done():
			return
		case col := <-e.col:
			//TODO: 输出日志等信息
			fmt.Printf("Run Task: %s\n", col.String())
		}
	}
}
