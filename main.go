package main

import (
	"context"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/handlers"
	"escort-book-user-log-consumer/repositories"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_SERVERS"),
		"group.id":           os.Getenv("KAFKA_GROUP_ID"),
		"auto.offset.reset":  "smallest",
		"enable.auto.commit": true,
	})
	consumer.SubscribeTopics(
		[]string{os.Getenv("REQUEST_LOG_TOPIC"), os.Getenv("CONNECTION_LOG_TOPIC")},
		nil,
	)
	handler := handlers.LogHandler{
		ConnectionLogRepository: &repositories.ConnectionLogRepository{
			Data: db.New(),
		},
		RequestLogRepository: &repositories.RequestLogRepository{
			Data: db.New(),
		},
		UserRepository: &repositories.UserRepository{
			Data: db.New(),
		},
	}
	run := true

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			handler.ProcessMessage(context.Background(), e)
			log.Println("PROCESSED MESSAGE")
		case kafka.PartitionEOF:
			log.Println("REACHED: ", e)
		case kafka.Error:
			log.Println("ERROR: ", e)
		default:
		}
	}

	consumer.Close()
}
