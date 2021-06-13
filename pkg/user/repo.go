package user

import "context"

type Repository interface {
	Migrate() error
	FindByID(context.Context, int, *Model) error
	FindByUsername(context.Context, string, *Model) error
}
