package repositories

import (
	"context"
	"testing"

	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"

	"github.com/stretchr/testify/assert"
)

func TestRequestLogRepositoryCreate(t *testing.T) {
	requestLogRepository := RequestLogRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := requestLogRepository.Create(ctxWithCancel, models.RequestLog{})
		assert.Error(t, err)
	})
}
