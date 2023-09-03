package eventhandler

import (
	"fire_agent/service"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Handlers struct {
	Service service.IService
	MyTopic string
}

func Handler(em *Handlers, client mqtt.Client) {
	log.Println("Topic", em.MyTopic)
	token := client.Subscribe("topic/fire_client", 0, em.Health)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token.WaitTimeout(time.Hour * 100)

	log.Println("Subscribed to topic eh ", em.MyTopic)
}

func (h *Handlers) Health(c mqtt.Client, m mqtt.Message) {
	log.Println("Received message Topic : ", m.Topic())
	log.Println("Data", m.Payload())
	h.Service.Health()
}
