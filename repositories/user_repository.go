package repositories

import (
	"context"

	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

//go:generate mockgen -destination=./mocks/iuser_repository.go -package=mocks --build_flags=--mod=mod . IUserRepository
type IUserRepository interface {
	GetByEmail(ctx context.Context, email string) (models.User, error)
}

type UserRepository struct {
	Data *db.PostgresClient
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT user_id, email FROM "user" WHERE email = $1;`
	row := r.Data.UserDB.QueryRowContext(ctx, query, email)

	var user models.User

	if err := row.Scan(&user.UserId, &user.Email); err != nil {
		return user, err
	}

	return user, nil
}
