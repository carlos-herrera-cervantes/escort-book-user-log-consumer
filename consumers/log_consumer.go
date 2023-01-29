package consumers

import (
	"context"

	"escort-book-user-log-consumer/config"
	"escort-book-user-log-consumer/handlers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/inconshreveable/log15"
)

type LogConsumer struct {
	EventHandler handlers.IEventHandler
}

var logger = log.New("consumers")

func (c LogConsumer) StartConsumer() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.InitializeKafka().BootstrapServers,
		"group.id":           config.InitializeKafka().GroupId,
		"auto.offset.reset":  "smallest",
		"enable.auto.commit": true,
	})
	topics := []string{
		config.InitializeKafka().Topics.ConnectionLog,
		config.InitializeKafka().Topics.RequestLog,
	}
	_ = consumer.SubscribeTopics(topics, nil)
	run := true

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			c.EventHandler.HandleEvent(context.Background(), e)
			log.Info("PROCESSED MESSAGE")
		case kafka.PartitionEOF:
			log.Info("REACHED: ", e)
		case kafka.Error:
			log.Error("ERROR: ", e)
			run = false
		}
	}

	_ = consumer.Close()
}
