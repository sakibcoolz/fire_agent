package eventhandler

import (
	"fire_agent/service"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Handlers struct {
	Service service.IService
	Topic   string
}

func Handler(em *Handlers, client mqtt.Client) {
	log.Println("Topic", em.Topic)
	token := client.Subscribe(em.Topic, 0, em.Health)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token.WaitTimeout(time.Hour * 100)

	log.Println("Subscribed to topic eh ", em.Topic)
}

func (h *Handlers) Health(c mqtt.Client, m mqtt.Message) {
	log.Println("Received message Topic : ", m.Topic())
	log.Println("Data", m.Payload())
	h.Service.Health()
}
