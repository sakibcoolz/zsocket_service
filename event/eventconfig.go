package event

import (
	"fmt"
	"time"
	"zsocket_service/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

var (
	log *zap.Logger
)

func init() {
	log = zap.NewExample()
}

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Info("Received message", zap.String("topic", msg.Topic()), zap.Any("data", msg.Payload()))
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Error("Lost Connection", zap.String("Error", err.Error()))
}

func GetMqttConfig(mqconfig config.Mqtt) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", mqconfig.HiveMQ))
	opts.SetClientID(mqconfig.ClientID)
	// opts.SetUsername(config.MQTT.Username)
	// opts.SetPassword(config.MQTT.Password)
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.OnConnect = ConnectHandler
	opts.OnConnectionLost = ConnectLostHandler
	opts.SetAutoReconnect(true)
	opts.SetKeepAlive(18 * time.Hour)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Fatal on GetMqttConfig", zap.Error(token.Error()))
	}

	return client
}
