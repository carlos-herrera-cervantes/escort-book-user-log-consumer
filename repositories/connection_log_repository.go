package repositories

import (
	"context"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

type IConnectionLogRepository interface {
	Create(ctx context.Context, connectionLog models.ConnectionLog) error
}

type ConnectionLogRepository struct {
	Data *db.Data
}

func (r *ConnectionLogRepository) Create(ctx context.Context, connectionLog models.ConnectionLog) error {
	query := "INSERT INTO connection_log VALUES ($1, $2, $3);"
	connectionLog.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		connectionLog.Id,
		connectionLog.UserId,
		connectionLog.LastConnection,
	)

	if err != nil {
		return err
	}

	return nil
}
