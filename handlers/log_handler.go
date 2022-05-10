package handlers

import (
	"context"
	"encoding/json"
	"escort-book-user-log-consumer/models"
	"escort-book-user-log-consumer/repositories"
	"escort-book-user-log-consumer/types"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type LogHandler struct {
	ConnectionLogRepository repositories.IConnectionLogRepository
	RequestLogRepository    repositories.IRequestLogRepository
	UserRepository          repositories.IUserRepository
}

func (h *LogHandler) ProcessMessage(ctx context.Context, message *kafka.Message) {
	topic := message.TopicPartition.Topic
	value := message.Value

	if *(topic) == "user-request" {
		var request types.RequestEvent
		json.Unmarshal(value, &request)
		requestLog := models.RequestLog{
			UserId:    request.UserId,
			Component: request.Component,
			Path:      request.Path,
			Method:    request.Method,
			Payload:   request.Payload,
		}

		err := h.RequestLogRepository.Create(ctx, requestLog)

		if err != nil {
			log.Println("WE GOT AN ERROR TRYING TO CREATE A REQUEST LOG: ", err.Error())
		}
		return
	}

	var lastConnection types.ConnectionEvent
	json.Unmarshal(value, &lastConnection)
	user, err := h.UserRepository.GetByEmail(ctx, lastConnection.Email)

	if err != nil {
		log.Println("WE GOT AN ERROR TRYING TO GET A USER: ", err.Error())
		return
	}

	connectionLog := models.ConnectionLog{
		LastConnection: time.Now(),
		UserId:         user.UserId,
	}

	err = h.ConnectionLogRepository.Create(ctx, connectionLog)

	if err != nil {
		log.Println("WE GOT AN ERROR TRYING TO CREATE A CONNECTION LOG: ", err.Error())
	}
}
