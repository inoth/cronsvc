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

	taskCount int

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
	return e.taskCount
}

func (e *Executor) AddTask(taskBody TaskBody) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.tasks[taskBody.ID]; !ok {

		val, ok := task.TaskMap[taskBody.Tag]
		if !ok {
			return fmt.Errorf("invalid task tag = %s", taskBody.Tag)
		}
		runID, err := e.c.AddFunc(taskBody.Crontab, func() {
			val().Run(e.col, taskBody.ID, taskBody.Body)
		})
		if err != nil {
			return errors.Wrap(err, "append task job err")
		}
		taskBody.runID = int(runID)

		taskBody.ctx, taskBody.cancel = context.WithCancel(e.ctx)
		e.tasks[taskBody.ID] = taskBody
		e.taskCount += 1

		fmt.Printf("start task %s success; current task %d\n", taskBody.ID, e.taskCount)
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
		e.taskCount -= 1

		fmt.Printf("remove task %s success; current task %d\n", taskId, e.taskCount)
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
	if err := e.c.Stop().Err(); err != nil && err != context.Canceled {
		return err
	}
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
