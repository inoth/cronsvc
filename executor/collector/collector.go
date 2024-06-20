package collector

import (
	"encoding/json"
	"time"
)

const (
	TyptRunningTime = "runtime"
)

type Collector struct {
	Call      string            `json:"call"`
	ID        string            `json:"id"`
	Field     map[string]string `json:"field"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
}

func NewCollectorWithRunning(id string, startTime time.Time, field map[string]string) Collector {
	return NewCollector(TyptRunningTime, id, startTime, field)
}

func NewCollector(call, id string, startTime time.Time, field map[string]string) Collector {
	return Collector{
		Call:      call,
		ID:        id,
		Field:     field,
		StartTime: startTime,
		EndTime:   time.Now(),
	}
}

func (tb *Collector) String() string {
	buf, _ := json.Marshal(tb)
	return string(buf)
}
