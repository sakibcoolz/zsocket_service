package eventhandler

import (
	"encoding/json"
	"time"
	"zsocket_service/model"
	"zsocket_service/service"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Event struct {
	service service.IService
	log     *zap.Logger
	client  mqtt.Client
	topic   string
}

func NewEvent(service service.IService, log *zap.Logger, client mqtt.Client, topic string) *Event {
	return &Event{
		service: service,
		log:     log,
		client:  client,
		topic:   topic,
	}
}

func (e *Event) Receiver() {
	token := e.client.Subscribe(e.topic, 0, e.Received)
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token.WaitTimeout(time.Millisecond * 100)

	e.log.Info("Subscribed to topic :", zap.String("Topic", e.topic))
}

func (e *Event) Received(client mqtt.Client, msg mqtt.Message) {
	e.log.Info("Executed")
	var data model.Msg
	json.Unmarshal(msg.Payload(), &data)
	e.service.MsgFiner(data)
}
