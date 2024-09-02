package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/ginsvr"
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (tc *TaskController) Prefix() string {
	return "/api"
}

func (tc *TaskController) Middlewares() []gin.HandlerFunc {
	return nil
}

func (c *TaskController) Routers() []ginsvr.Router {
	return []ginsvr.Router{}
}
