package httpcol

import (
	"github.com/inoth/cronsvc/executor/collector"
	"github.com/inoth/cronsvc/executor/task"
)

const tag = "http_collector"

type HttpCollector struct{}

func (hc *HttpCollector) Run(col chan<- collector.Collector, taskId string, args map[string]string) {
	col <- collector.NewCollector(taskId, args)
}

func init() {
	task.AddCronTask(tag, func() task.CornTask {
		return &HttpCollector{}
	})
}
