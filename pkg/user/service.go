package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Authenticate(ctx context.Context, string, password string, user *Model) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) Authenticate(ctx context.Context, username string, password string, user *Model) error {
	err := service.repo.FindByUsername(ctx, username, user)
	if err != nil {
		return err //TODO: custom type for errors
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(user.PasswordHash),
	); err != nil {
		return err
	}

	return nil
}
