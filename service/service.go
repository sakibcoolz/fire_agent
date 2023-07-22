package service

import (
	"encoding/json"
	"fire_agent/model"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

type Service struct {
	DB      *gorm.DB
	Event   mqtt.Client
	Topic   string
	MyTopic string
}

type IService interface {
	Health()
}

func (s *Service) Health() {
	payload := model.Msg{Topic: s.MyTopic,
		Format: "Talks",
		Talks: []*model.Talks{
			&model.Talks{Name: "Sakib",
				Text: "Hi"},
		},
	}

	str, _ := json.Marshal(payload)

	log.Println(string(str))

	log.Println(s.Topic)
	token := s.Event.Publish(s.Topic, 0, false, str)
	if token.Error() != nil {
		log.Println(token.Error())
	}
	token.Wait()
	time.Sleep(1 * time.Second)
}
