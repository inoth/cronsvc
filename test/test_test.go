package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/inoth/cronsvc"
	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/executor"
	httpapi "github.com/inoth/cronsvc/http-api"
	"github.com/inoth/cronsvc/internal/util"
)

func TestNewCronSvc(t *testing.T) {
	c := cronsvc.New(
		cronsvc.WithConfig(config.NewConfig()),
		cronsvc.WithServer(
			executor.New(),
			httpapi.New(httpapi.WithGET("", func(c *gin.Context) {
				executor.ReceiverTask(executor.TaskBody{
					ID:      util.UUID(),
					Title:   "test task",
					Status:  1,
					Tag:     util.UUID(),
					Crontab: "* 5 4 * * sun",
					Body: map[string]string{
						"val1": "aaaa",
						"val2": "bbbb",
						"val3": "cccc",
					},
				})
				c.String(200, " hello world ")
			})),
		),
	)
	if err := c.Run(); err != nil {
		t.Logf(err.Error())
	}
}
