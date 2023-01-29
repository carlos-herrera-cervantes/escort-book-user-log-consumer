package config

import "os"

type kafka struct {
	BootstrapServers string
	GroupId          string
	Topics           topic
}

type topic struct {
	RequestLog    string
	ConnectionLog string
}

var singleKafka *kafka

func InitializeKafka() *kafka {
	if singleKafka != nil {
		return singleKafka
	}

	lock.Lock()
	defer lock.Unlock()

	singleKafka = &kafka{
		BootstrapServers: os.Getenv("KAFKA_SERVERS"),
		GroupId:          os.Getenv("KAFKA_GROUP_ID"),
		Topics: topic{
			RequestLog:    "user-request",
			ConnectionLog: "user-last-connection",
		},
	}

	return singleKafka
}
