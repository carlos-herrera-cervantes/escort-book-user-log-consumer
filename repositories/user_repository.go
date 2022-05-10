package repositories

import (
	"context"
	"escort-book-user-log-consumer/db"
	"escort-book-user-log-consumer/models"
)

type IUserRepository interface {
	GetByEmail(ctx context.Context, email string) (models.User, error)
}

type UserRepository struct {
	Data *db.Data
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT user_id, email FROM "user" WHERE email = $1;`
	row := r.Data.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.UserId, &user.Email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
