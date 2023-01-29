package repositories

import (
	"context"

	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

//go:generate mockgen -destination=./mocks/iconnection_log_repository.go -package=mocks --build_flags=--mod=mod . IConnectionLogRepository
type IConnectionLogRepository interface {
	GetByUserId(ctx context.Context, userId string) (models.ConnectionLog, error)
	Create(ctx context.Context, connectionLog models.ConnectionLog) error
	UpdateById(ctx context.Context, id string, connectionLog models.ConnectionLog) error
}

type ConnectionLogRepository struct {
	Data *db.PostgresClient
}

func (r *ConnectionLogRepository) GetByUserId(ctx context.Context, userId string) (models.ConnectionLog, error) {
	query := "SELECT * from connection_log WHERE user_id = $1;"
	row := r.Data.UserDB.QueryRowContext(ctx, query, userId)

	var connectionLog models.ConnectionLog

	if err := row.Scan(&connectionLog.Id, &connectionLog.UserId, &connectionLog.LastConnection); err != nil {
		return connectionLog, err
	}

	return connectionLog, nil
}

func (r *ConnectionLogRepository) Create(ctx context.Context, connectionLog models.ConnectionLog) error {
	query := "INSERT INTO connection_log VALUES ($1, $2, $3);"
	connectionLog.SetDefaultValues()

	if _, err := r.Data.UserDB.ExecContext(
		ctx,
		query,
		connectionLog.Id,
		connectionLog.UserId,
		connectionLog.LastConnection,
	); err != nil {
		return err
	}

	return nil
}

func (r *ConnectionLogRepository) UpdateById(ctx context.Context, id string, connectionLog models.ConnectionLog) error {
	query := "UPDATE connection_log SET last_connection  = $1 WHERE id = $2;"

	if _, err := r.Data.UserDB.ExecContext(ctx, query, connectionLog.LastConnection, connectionLog.Id); err != nil {
		return err
	}

	return nil
}
