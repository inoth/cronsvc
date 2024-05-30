package executor

import (
	"github.com/inoth/cronsvc/logger"
	"github.com/inoth/cronsvc/registry"
)

type Option func(*option)

type option struct {
	id      string
	name    string
	version string

	reg registry.Register
	log logger.Logger
}

func Name(name string) Option {
	return func(o *option) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *option) {
		o.version = version
	}
}

func Registrar(r registry.Register) Option {
	return func(o *option) {
		o.reg = r
	}
}

func Logger(log logger.Logger) Option {
	return func(o *option) {
		o.log = log
	}
}
