package executor

import (
	"context"
	"fmt"
)

const (
	name = "executor"
)

// 启动cron监听进程
type Executor struct {
	MaxTask int `toml:"max_task"`
}

func (e *Executor) Name() string {
	return name
}

func (e *Executor) Start(ctx context.Context) error {
	fmt.Printf("%s start run, max_task %d", name, e.MaxTask)
	return nil
}

func (e *Executor) Stop(ctx context.Context) error {
	return nil
}
