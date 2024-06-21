package test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/inoth/cronsvc"
	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/executor"
	httpapi "github.com/inoth/cronsvc/http-api"
	"github.com/inoth/cronsvc/internal/util"
	"github.com/inoth/cronsvc/metric"
	"github.com/inoth/cronsvc/perf"

	_ "github.com/inoth/cronsvc/executor/task/all"
)

func TestNewCronSvc(t *testing.T) {
	c := cronsvc.New(
		cronsvc.WithConfig(config.NewConfig()),
		cronsvc.WithServer(
			perf.New(perf.WithPort(":9090")),
			metric.New(),
			executor.New(),
			httpapi.New(
				httpapi.WithGET("", func(c *gin.Context) {
					id := util.UUID()
					executor.ReceiverTask(executor.TaskBody{
						ID:      id,
						Title:   "test task",
						Status:  1,
						Tag:     "http_collector",
						Crontab: "0/5 * * * * *",
						Body: map[string]string{
							"val1": "aaaa",
							"val2": "bbbb",
							"val3": "cccc",
						},
					})
					c.String(200, fmt.Sprintf("task id: %s", id))
				}),
				httpapi.WithGET("/del/:id", func(c *gin.Context) {
					id := c.Param("id")
					executor.RemoveTask(id)
					c.String(200, "ok")
				})),
		),
	)
	if err := c.Run(); err != nil {
		t.Logf(err.Error())
	}
}
