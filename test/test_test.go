package test

import (
	"testing"

	"github.com/inoth/cronsvc"
	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/executor"
)

func TestNewCronSvc(t *testing.T) {
	c := cronsvc.New(
		cronsvc.WithConfig(config.NewConfig()),
		cronsvc.WithServer(&executor.Executor{}),
	)
	if err := c.Run(); err != nil {
		t.Logf(err.Error())
	}
}
