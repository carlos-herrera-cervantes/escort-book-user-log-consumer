package repositories

import (
	"context"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

type IConnectionLogRepository interface {
	GetByUserId(ctx context.Context, userId string) (models.ConnectionLog, error)
	Create(ctx context.Context, connectionLog models.ConnectionLog) error
	UpdateById(ctx context.Context, id string, connectionLog models.ConnectionLog) error
}

type ConnectionLogRepository struct {
	Data *db.Data
}

func (r *ConnectionLogRepository) GetByUserId(ctx context.Context, userId string) (models.ConnectionLog, error) {
	query := "SELECT * from connection_log WHERE user_id = $1;"
	row := r.Data.DB.QueryRowContext(ctx, query, userId)

	var connectionLog models.ConnectionLog
	err := row.Scan(&connectionLog.Id, &connectionLog.UserId, &connectionLog.LastConnection)

	if err != nil {
		return models.ConnectionLog{}, err
	}

	return connectionLog, nil
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

func (r *ConnectionLogRepository) UpdateById(ctx context.Context, id string, connectionLog models.ConnectionLog) error {
	query := "UPDATE connection_log SET last_connection  = $1 WHERE id = $2;"
	_, err := r.Data.DB.ExecContext(ctx, query, connectionLog.LastConnection, connectionLog.Id)

	if err != nil {
		return err
	}

	return nil
}
