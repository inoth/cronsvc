package collector

import (
	"encoding/json"
	"time"
)

const (
	TyptRunningTime = "runtime"
)

type RunningTimeFunc func(col chan<- Collector)

type Collector struct {
	Call      string            `json:"call"`
	ID        string            `json:"id"`
	Field     map[string]string `json:"field"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
}

func (tb *Collector) String() string {
	buf, _ := json.Marshal(tb)
	return string(buf)
}

func NewCollectorWithRunning(id string, startTime time.Time, field map[string]string) Collector {
	return newCollector(TyptRunningTime, id, startTime, field)
}

func NewCollector(id string, startTime time.Time, field map[string]string) Collector {
	return newCollector("", id, startTime, field)
}

func newCollector(call, id string, startTime time.Time, field map[string]string) Collector {
	return Collector{
		Call:      call,
		ID:        id,
		Field:     field,
		StartTime: startTime,
		EndTime:   time.Now(),
	}
}

func RunningTime(taskId string) RunningTimeFunc {
	start := time.Now()
	return func(col chan<- Collector) {
		col <- NewCollectorWithRunning(taskId, start, nil)
	}
}
