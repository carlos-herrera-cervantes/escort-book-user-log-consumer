package handlers

import (
	"context"
	"errors"
	"testing"

	"escort-book-user-log-consumer/config"
	"escort-book-user-log-consumer/models"
	"escort-book-user-log-consumer/repositories/mocks"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
)

func TestLogHandlerHandleEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConnectionLogRepository := mocks.NewMockIConnectionLogRepository(ctrl)
	mockRequestLogRepository := mocks.NewMockIRequestLogRepository(ctrl)
	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	logHandler := LogHandler{
		ConnectionLogRepository: mockConnectionLogRepository,
		RequestLogRepository:    mockRequestLogRepository,
		UserRepository:          mockUserRepository,
	}

	t.Run("Should exit when creating a request log fails", func(t *testing.T) {
		mockRequestLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockUserRepository.
			EXPECT().
			GetByEmail(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			GetByUserId(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			UpdateById(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		topic := config.InitializeKafka().Topics.RequestLog
		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic,
			},
			Value: []byte(`{
				"userId": "63d441eb995ce0e571b307e0",
				"component": "test",
				"path": "/authorizer/login",
				"method": "POST",
				"payload": "{}",
			}`),
		}

		logHandler.HandleEvent(context.Background(), &message)
	})

	t.Run("Should return error when getting a user fails", func(t *testing.T) {
		mockRequestLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockUserRepository.
			EXPECT().
			GetByEmail(gomock.Any(), gomock.Any()).
			Return(models.User{}, errors.New("dummy error")).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			GetByUserId(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			UpdateById(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		topic := config.InitializeKafka().Topics.ConnectionLog
		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic,
			},
			Value: []byte(`{"email": "test@example.com"}`),
		}

		logHandler.HandleEvent(context.Background(), &message)
	})

	t.Run("Should return error when creating connection log fails", func(t *testing.T) {
		mockRequestLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockUserRepository.
			EXPECT().
			GetByEmail(gomock.Any(), gomock.Any()).
			Return(models.User{UserId: "63d4496736b36b12da077c77"}, nil).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			GetByUserId(gomock.Any(), gomock.Any()).
			Return(models.ConnectionLog{}, errors.New("dummy error")).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			UpdateById(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		topic := config.InitializeKafka().Topics.ConnectionLog
		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic,
			},
			Value: []byte(`{"email": "test@example.com"}`),
		}

		logHandler.HandleEvent(context.Background(), &message)
	})

	t.Run("Should log error when update connection log fails", func(t *testing.T) {
		mockRequestLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockUserRepository.
			EXPECT().
			GetByEmail(gomock.Any(), gomock.Any()).
			Return(models.User{UserId: "63d4496736b36b12da077c77"}, nil).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			GetByUserId(gomock.Any(), gomock.Any()).
			Return(models.ConnectionLog{
				Id: "63d55228529dc4d065f064f8",
			}, nil).
			Times(1)
		mockConnectionLogRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)
		mockConnectionLogRepository.
			EXPECT().
			UpdateById(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		topic := config.InitializeKafka().Topics.ConnectionLog
		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &topic,
			},
			Value: []byte(`{"email": "test@example.com"}`),
		}

		logHandler.HandleEvent(context.Background(), &message)
	})
}
