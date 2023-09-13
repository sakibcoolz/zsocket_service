package service

import (
	"encoding/json"
	"time"
	"zsocket_service/model"

	"go.uber.org/zap"
)

func (s *Service) MsgFiner(msg model.Msg) {
	switch msg.Format {
	case "Talks":
		s.log.Info("Message", zap.Any("Data", msg.Talks))
		s.Talks(msg)
	case "Commands":
		s.log.Info("Message", zap.Any("Data", msg.Commands))
	case "Databyte":
		s.log.Info("Message", zap.Any("Data", msg.Databyte))
	}
}

func (s *Service) Talks(msg model.Msg) {
	talks := msg.Talks
	for _, value := range talks {
		if value.Text == "Hi" {
			data := model.Talks{
				Name: msg.Topic,
				Text: "Hi i am server",
			}

			body, _ := json.Marshal(data)

			token := s.client.Publish(msg.Topic, 0, false, body)
			if token.Error() != nil {
				s.log.Error("Cant sned ", zap.String("Topic", msg.Topic), zap.Error(token.Error()))
			}
			token.Wait()
			time.Sleep(1 * time.Second)
		}
	}
}
