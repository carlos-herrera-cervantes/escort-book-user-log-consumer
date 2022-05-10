package repositories

import (
	"context"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

type IRequestLogRepository interface {
	Create(ctx context.Context, requestLog models.RequestLog) error
}

type RequestLogRepository struct {
	Data *db.Data
}

func (r *RequestLogRepository) Create(ctx context.Context, requestLog models.RequestLog) error {
	query := "INSERT INTO request_log VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	requestLog.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
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
	)

	if err != nil {
		return err
	}

	return nil
}
