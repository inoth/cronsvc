package collector

import (
	"encoding/json"
	"time"
)

type Collector struct {
	ID       string            `json:"id"`
	Field    map[string]string `json:"field"`
	Timespan time.Time         `json:"timespan"`
}

func NewCollector(id string, field map[string]string) Collector {
	return Collector{
		ID:       id,
		Field:    field,
		Timespan: time.Now(),
	}
}

func (tb *Collector) String() string {
	buf, _ := json.Marshal(tb)
	return string(buf)
}
