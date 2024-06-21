package perf

import (
	"context"
	"net/http"
	_ "net/http/pprof"

	"github.com/pkg/errors"
)

const (
	name = "perf"
)

type Perf struct {
	option
}

func New(opts ...Option) *Perf {
	o := option{
		Port: ":9059",
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &Perf{
		option: o,
	}
}

func (p *Perf) Name() string {
	return name
}

func (p *Perf) Start(ctx context.Context) error {
	if err := http.ListenAndServe(p.Port, nil); err != nil {
		return errors.Wrap(err, "start pprof err")
	}
	return nil
}

func (p *Perf) Stop(ctx context.Context) error {
	return nil
}
