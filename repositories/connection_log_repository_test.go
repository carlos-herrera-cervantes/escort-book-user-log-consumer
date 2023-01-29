package repositories

import (
	"context"
	"testing"

	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"

	"github.com/stretchr/testify/assert"
)

func TestConnectionLogRepositoryGetByUserId(t *testing.T) {
	connectionLogRepository := ConnectionLogRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := connectionLogRepository.GetByUserId(context.Background(), "63d554c99970a6d0f6953aab")
		assert.Error(t, err)
	})
}

func TestConnectionLogRepositoryCreate(t *testing.T) {
	connectionLogRepository := ConnectionLogRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := connectionLogRepository.Create(ctxWithCancel, models.ConnectionLog{})
		assert.Error(t, err)
	})
}

func TestConnectionLogRepositoryUpdateById(t *testing.T) {
	connectionLogRepository := ConnectionLogRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()

		err := connectionLogRepository.UpdateById(
			ctxWithCancel,
			"63d554c99970a6d0f6953aab",
			models.ConnectionLog{},
		)
		assert.Error(t, err)
	})
}
