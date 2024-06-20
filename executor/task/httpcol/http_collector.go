package httpcol

import (
	"context"
	"fmt"
	"time"

	"github.com/inoth/cronsvc/executor/collector"
	"github.com/inoth/cronsvc/executor/task"
)

const tag = "http_collector"

type HttpCollector struct{}

func (hc *HttpCollector) Run(ctx context.Context, col chan<- collector.Collector, args map[string]string) {
	taskId := ctx.Value("taskId").(string)
	defer runningTime(taskId, col)()

	fmt.Printf("taskid = %s; http_collector run dosomething %+v", taskId, args)
	time.Sleep(time.Microsecond * 100)
}

func init() {
	task.AddCronTask(tag, func() task.CornTask {
		return &HttpCollector{}
	})
}

func runningTime(taskId string, col chan<- collector.Collector) func() {
	start := time.Now()
	return func() {
		col <- collector.NewCollectorWithRunning(taskId, start, nil)
	}
}
