package handlers

import (
	"context"
	"encoding/json"
	"time"

	"escort-book-user-log-consumer/config"
	"escort-book-user-log-consumer/models"
	"escort-book-user-log-consumer/repositories"
	"escort-book-user-log-consumer/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/inconshreveable/log15"
)

type LogHandler struct {
	ConnectionLogRepository repositories.IConnectionLogRepository
	RequestLogRepository    repositories.IRequestLogRepository
	UserRepository          repositories.IUserRepository
}

var logger = log.New("handlers")

func (h *LogHandler) HandleEvent(ctx context.Context, message *kafka.Message) {
	topic := message.TopicPartition.Topic
	value := message.Value

	if *(topic) == config.InitializeKafka().Topics.RequestLog {
		var request types.RequestEvent
		_ = json.Unmarshal(value, &request)
		requestLog := models.RequestLog{
			UserId:    request.UserId,
			Component: request.Component,
			Path:      request.Path,
			Method:    request.Method,
			Payload:   request.Payload,
		}

		if err := h.RequestLogRepository.Create(ctx, requestLog); err != nil {
			log.Error("WE GOT AN ERROR TRYING TO CREATE A REQUEST LOG: ", err.Error())
		}

		return
	}

	var lastConnection types.ConnectionEvent
	_ = json.Unmarshal(value, &lastConnection)
	user, err := h.UserRepository.GetByEmail(ctx, lastConnection.Email)

	if err != nil {
		log.Error("WE GOT AN ERROR TRYING TO GET A USER: ", err.Error())
		return
	}

	existConnectionLog, err := h.ConnectionLogRepository.GetByUserId(ctx, user.UserId)

	if err != nil {
		connectionLog := models.ConnectionLog{
			LastConnection: time.Now(),
			UserId:         user.UserId,
		}

		if err = h.ConnectionLogRepository.Create(ctx, connectionLog); err != nil {
			log.Error("WE GOT AN ERROR TRYING TO CREATE A CONNECTION LOG: ", err.Error())
		}

		return
	}

	existConnectionLog.LastConnection = time.Now()

	if err = h.ConnectionLogRepository.UpdateById(
		ctx,
		existConnectionLog.Id,
		existConnectionLog,
	); err != nil {
		log.Error("WE GOT AN ERROR TRYING TO UPDATE A CONNECTION LOG: ", err.Error())
	}
}
