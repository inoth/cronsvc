package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/inoth/cronsvc"
	"github.com/inoth/cronsvc/config"
	"github.com/inoth/cronsvc/executor"
	httpapi "github.com/inoth/cronsvc/http-api"
	"github.com/inoth/cronsvc/internal/util"
	"github.com/inoth/cronsvc/metric"
)

type TaskRequest struct {
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Tag     string            `json:"tag"`
	Crontab string            `json:"crontab"`
	Body    map[string]string `json:"body"`
}

func Main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runApp() {
	c := cronsvc.New(
		cronsvc.WithConfig(config.NewConfig()),
		cronsvc.WithServer(
			// perf.New(),
			metric.New(),
			executor.New(),
			httpapi.New(
				httpapi.WithPOST("/add", func(c *gin.Context) {
					var req TaskRequest
					if err := c.ShouldBindJSON(&req); err != nil {
						c.String(http.StatusOK, err.Error())
						return
					}
					id := util.UUID()
					executor.ReceiverTask(executor.TaskBody{
						ID:      id,
						Title:   req.Title,
						Tag:     req.Tag,
						Crontab: req.Crontab,
						Body:    req.Body,
					})
					c.String(http.StatusOK, id)
				}),
				httpapi.WithPOST("/del/:id", func(c *gin.Context) {
					id := c.Param("id")
					executor.RemoveTask(id)
					c.String(http.StatusOK, "ok")
				}),
			),
		),
	)
	if err := c.Run(); err != nil && err != context.Canceled {
		fmt.Printf("run cronsvc err %v\n", err)
	}
}
