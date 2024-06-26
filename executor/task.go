package executor

import (
	"context"
	"encoding/json"
)

type TaskBody struct {
	ctx    context.Context
	cancel func()

	runID   int               `json:"-"`
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Tag     string            `json:"tag"`
	Crontab string            `json:"crontab"`
	Body    map[string]string `json:"body"`
}

func (tb *TaskBody) String() string {
	buf, _ := json.Marshal(tb)
	return string(buf)
}

func (tb *TaskBody) Equal(o any) bool {
	if tb == nil && o == nil {
		return true
	}
	if tb == nil || o == nil {
		return false
	}

	t, ok := o.(*TaskBody)
	if !ok {
		return false
	}

	return tb.ID == t.ID &&
		tb.Title == t.Title &&
		tb.Tag == t.Tag &&
		tb.Crontab == t.Crontab
}

func ReceiverTask(msg TaskBody) {
	ecr.receiver <- msg
}

func RemoveTask(taskId string) {
	ecr.RemoveTask(taskId)
}
