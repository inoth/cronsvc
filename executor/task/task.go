package task

import "github.com/inoth/cronsvc/executor/collector"

type CornTask interface {
	Run(col chan<- collector.Collector, taskId string, args map[string]string)
}

type Creator func() CornTask

var TaskMap = map[string]Creator{}

func AddCronTask(name string, creator Creator) {
	TaskMap[name] = creator
}
