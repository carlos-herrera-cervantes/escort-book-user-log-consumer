package repositories

import (
	"context"
	"testing"

	"escort-book-user-log-consumer/db"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryGetByEmail(t *testing.T) {
	userRepository := UserRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := userRepository.GetByEmail(context.Background(), "bad@example.com")
		assert.Error(t, err)
	})
}
