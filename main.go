package main

import (
	"escort-book-user-log-consumer/consumers"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/handlers"
	"escort-book-user-log-consumer/repositories"
)

func main() {
	logHandler := &handlers.LogHandler{
		ConnectionLogRepository: &repositories.ConnectionLogRepository{
			Data: db.NewPostgresClient(),
		},
		RequestLogRepository: &repositories.RequestLogRepository{
			Data: db.NewPostgresClient(),
		},
		UserRepository: &repositories.UserRepository{
			Data: db.NewPostgresClient(),
		},
	}
	logConsumer := consumers.LogConsumer{
		EventHandler: logHandler,
	}

	logConsumer.StartConsumer()
}
