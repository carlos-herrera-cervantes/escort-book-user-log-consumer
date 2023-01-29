package repositories

import (
	"context"

	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

//go:generate mockgen -destination=./mocks/irequest_log_repository.go -package=mocks --build_flags=--mod=mod . IRequestLogRepository
type IRequestLogRepository interface {
	Create(ctx context.Context, requestLog models.RequestLog) error
}

type RequestLogRepository struct {
	Data *db.PostgresClient
}

func (r *RequestLogRepository) Create(ctx context.Context, requestLog models.RequestLog) error {
	query := "INSERT INTO request_log VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	requestLog.SetDefaultValues()

	if _, err := r.Data.UserDB.ExecContext(
		ctx,
		query,
		requestLog.Id,
		requestLog.UserId,
		requestLog.Component,
		requestLog.Path,
		requestLog.Method,
		requestLog.Payload,
		requestLog.CreatedAt,
		requestLog.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}
