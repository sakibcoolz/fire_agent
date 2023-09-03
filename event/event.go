package event

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Println("Received message Topic", msg.Topic(), "Data:", msg.Payload())
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Lost Connection Error : ", err.Error())
}

func GetMqttConfig() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", os.Getenv("LIVE")))
	opts.SetClientID("fire_client")
	// opts.SetUsername(config.MQTT.Username)
	// opts.SetPassword(config.MQTT.Password)
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.OnConnect = ConnectHandler
	opts.OnConnectionLost = ConnectLostHandler
	opts.SetAutoReconnect(true)
	opts.SetKeepAlive(18 * time.Hour)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Fatal on GetMqttConfig", token.Error())
	}

	return client
}
