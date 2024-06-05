package cronsvc

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/internal/util"
	"golang.org/x/sync/errgroup"
)

type CronExecutor struct {
	opt option
	// mu     sync.Mutex
	ctx    context.Context
	cancel func()
}

func New(opts ...Option) *CronExecutor {
	o := option{
		id:      util.UUID(),
		version: util.UUID(),
		ctx:     context.Background(),
		sigs:    []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &CronExecutor{
		opt:    o,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (ce *CronExecutor) ID() string      { return ce.opt.id }
func (ce *CronExecutor) Name() string    { return ce.opt.name }
func (ce *CronExecutor) Version() string { return ce.opt.version }

func (ce *CronExecutor) Run() (err error) {

	if ce.opt.cfg == nil {
		return ErrNotConfig
	}

	wg := sync.WaitGroup{}
	eg, ctx := errgroup.WithContext(ce.ctx)

	for _, svc := range ce.opt.svcs {
		svc := svc
		if err := ce.opt.cfg.PrimitiveDecode(svc.(config.ConfigureMatcher)); err != nil {
			return err
		}
		eg.Go(func() error {
			<-ctx.Done()
			return svc.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return svc.Start(ctx)
		})
	}

	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, ce.opt.sigs...)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
			return ce.Stop()
		}
	})
	if err = eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (ce *CronExecutor) Stop() error {
	if ce.cancel != nil {
		ce.cancel()
	}
	return nil
}