package controller

import (
	"fire_agent/config"
	"fire_agent/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Config  *config.AgentConfig
	Service service.IService
}

func (c Controller) Health(ctx *gin.Context) {
	c.Service.Health()
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
