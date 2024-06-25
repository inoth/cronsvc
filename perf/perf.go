package perf

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/pkg/errors"
)

const (
	name = "perf"
)

type Perf struct {
	option

	svr *http.Server
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
	p.svr = &http.Server{
		Addr: p.Port,
	}
	if err := p.svr.ListenAndServe(); err != nil {
		return errors.Wrap(err, "start pprof err")
	}
	return nil
}

func (p *Perf) Stop(ctx context.Context) error {
	fmt.Println("Done perf..............")
	return p.svr.Shutdown(ctx)
}
