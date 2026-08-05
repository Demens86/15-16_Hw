package users

import (
	"context"
	"prac/models"
)

func (r *Repo) CreateUser(ctx context.Context, name, email string, age int, is_active bool) (*models.Users, error) {
	newUser := new(models.Users)

	err := r.db.GetContext(ctx, newUser,
		`insert into users (name, email, age, is_active)
values ($1, $2, $3, $4) returning *,
name, email, age, is_active`)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
