package app

import (
	"fire_agent/config"
	"fire_agent/controller"
	"fire_agent/event"
	"fire_agent/eventhandler"
	"fire_agent/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UrlMapping(r *gin.Engine, configs *config.AgentConfig) *gin.Engine {

	log.Println(configs)

	client := event.GetMqttConfig()

	service := &service.Service{
		DB:      &gorm.DB{},
		Event:   client,
		Topic:   configs.Topic,
		MyTopic: fmt.Sprintf("%s_%s", config.GetHostName(), os.Getenv("KEY")),
	}

	handlers := eventhandler.Handlers{
		Service: service,
		MyTopic: fmt.Sprintf("%s_%s", config.GetHostName(), os.Getenv("KEY")),
	}

	controller := &controller.Controller{
		Config:  configs,
		Service: service,
	}

	go eventhandler.Handler(&handlers, client)

	r.GET("/health", controller.Health)

	return r
}
