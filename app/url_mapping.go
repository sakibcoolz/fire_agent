package app

import (
	"fire_agent/config"
	"fire_agent/controller"
	"fire_agent/event"
	"fire_agent/eventhandler"
	"fire_agent/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UrlMapping(r *gin.Engine, configs *config.LoginResponse) *gin.Engine {

	log.Println(configs)

	client := event.GetMqttConfig(configs)

	service := &service.Service{
		DB:    &gorm.DB{},
		Event: client,
		Topic: configs.Topic,
	}

	handlers := eventhandler.Handlers{
		Service: service,
		Topic:   configs.Topic,
	}

	controller := &controller.Controller{
		Config:  configs,
		Service: service,
	}

	go eventhandler.Handler(&handlers, client)

	r.GET("/health", controller.Health)

	return r
}
