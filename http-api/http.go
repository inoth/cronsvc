package httpapi

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// 开启任务添加服务，通过channel发送task到executor
const (
	name = "gin"
)

type CronHttpServer struct {
	option
	sfg singleflight.Group
}

func New(opts ...Option) *CronHttpServer {
	o := option{
		engine:  gin.New(),
		handles: make([]Handler, 0),
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &CronHttpServer{
		option: o,
		sfg:    singleflight.Group{},
	}
}

func (h *CronHttpServer) Name() string {
	return name
}

func (e *CronHttpServer) Start(ctx context.Context) error {

	e.loadRouter()

	svr := &http.Server{
		Addr:           e.Port,
		Handler:        e.engine,
		ReadTimeout:    time.Second * time.Duration(e.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(e.WriteTimeout),
		MaxHeaderBytes: 1 << uint(e.MaxHeaderBytes),
	}

	if e.TLS {
		svr.TLSConfig = &tls.Config{
			MinVersion:               tls.VersionTLS13,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			},
		}
		return svr.ListenAndServeTLS(e.Cert, e.Key)
	} else {
		return svr.ListenAndServe()
	}
}

func (e *CronHttpServer) Stop(ctx context.Context) error {
	return nil
}

func (e *CronHttpServer) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
	return e.sfg.Do(key, fn)
}

func (e *CronHttpServer) loadRouter() {
	for _, h := range e.handles {
		for _, r := range h.Routers() {
			e.engine.Handle(
				r.Method,
				h.Prefix()+"/"+r.Path,
				append(h.Middlewares(), r.Handle...)...,
			)
		}
	}
}
