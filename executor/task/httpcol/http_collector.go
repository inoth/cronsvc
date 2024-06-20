package httpcol

import (
	"context"
	"time"

	"github.com/inoth/cronsvc/executor/collector"
	"github.com/inoth/cronsvc/executor/task"
)

const tag = "http_collector"

type HttpCollector struct{}

func (hc *HttpCollector) Run(ctx context.Context, col chan<- collector.Collector, args map[string]string) {
	taskId := ctx.Value("taskId").(string)
	defer collector.RunningTime(taskId)(col)

	time.Sleep(time.Millisecond * 100)
}

func init() {
	task.AddCronTask(tag, func() task.CornTask {
		return &HttpCollector{}
	})
}
