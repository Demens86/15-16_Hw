package usecase

import (
	"context"
	"errors"
	"net/mail"
	"prac/models"
)

func (c *Case) CreateUser(ctx context.Context, userData *models.UserRequest) (*models.User, error) {
	user := &models.User{}
	isValidEmail := func(email string) bool {
		_, err := mail.ParseAddress(email)
		return err == nil
	}
	if isValidEmail(userData.Email) {
		user.Email = userData.Email
	} else {
		return nil, errors.New("invalid email")
	}
	user.Name = userData.Name
	user.Age = userData.Age
	user.IsActive = userData.IsActive

	err := c.Repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
