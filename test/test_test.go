package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/inoth/cronsvc"
	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/executor"
	httpapi "github.com/inoth/cronsvc/http-api"
)

func TestNewCronSvc(t *testing.T) {
	c := cronsvc.New(
		cronsvc.WithConfig(config.NewConfig()),
		cronsvc.WithServer(
			executor.New(),
			httpapi.New(httpapi.WithGET("", func(c *gin.Context) {
				executor.ReceiverTask([]byte("xxxx"))
				c.String(200, " hello world ")
			})),
		),
	)
	if err := c.Run(); err != nil {
		t.Logf(err.Error())
	}
}
