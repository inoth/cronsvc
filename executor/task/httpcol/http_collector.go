package httpcol

import (
	"context"

	"github.com/inoth/cronsvc/executor/collector"
	"github.com/inoth/cronsvc/executor/task"
)

const tag = "http_collector"

type HttpCollector struct{}

func (hc *HttpCollector) Run(ctx context.Context, col chan<- collector.Collector, args map[string]string) {
	taskId := ctx.Value("taskId").(string)
	col <- collector.NewCollector(taskId, args)
}

func init() {
	task.AddCronTask(tag, func() task.CornTask {
		return &HttpCollector{}
	})
}
