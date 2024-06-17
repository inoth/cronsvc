package task

import (
	"context"

	"github.com/inoth/cronsvc/executor/collector"
)

type CornTask interface {
	Run(ctx context.Context, col chan<- collector.Collector, args map[string]string)
}

type Creator func() CornTask

var TaskMap = map[string]Creator{}

func AddCronTask(name string, creator Creator) {
	TaskMap[name] = creator
}
