package repo

import (
	"context"
	"errors"
	"prac/models"

	"github.com/lib/pq"
)

func (r *Repo) CreateUser(ctx context.Context, user *models.User) error {

	err := r.db.GetContext(ctx, user,
		`insert into users (name, email, age, is_active)
	values ($1, $2, $3, $4) returning *`,
		user.Name, user.Email, user.Age, user.IsActive)
	if err != nil {
		if pqErr, ok := errors.AsType[*pq.Error](err); ok {
			if pqErr.Code == "23505" {
				return errors.New("email already exists")
			}
		}
	}

	return err
}
