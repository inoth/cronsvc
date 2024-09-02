package server

import (
	"github.com/inoth/cronsvc/internal/controller"
	"github.com/inoth/toybox/ginsvr"
)

func NewHttpGinServer(tc *controller.TaskController) *ginsvr.GinHttpServer {
	return ginsvr.New(
		ginsvr.WithHandlers(tc),
	)
}
