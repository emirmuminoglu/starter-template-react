package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type pgRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &pgRepository{db: db}
}

func (repo *pgRepository) Migrate() error {
	_, err := repo.db.Exec(context.Background(), `
		CREATE TABLE user (
			id integer PRIMARY KEY,
			username text,
			password varchar(60)
		)
	`)

	return err
}

func (repo *pgRepository) FindByID(ctx context.Context, id int, user *Model) error {
	return repo.db.QueryRow(ctx,
		`
			SELECT 
				id, username, password
			FROM
				user
			WHERE id = $1 
		`,
		id,
	).Scan(&user.ID, &user.Username, &user.PasswordHash)
}

func (repo *pgRepository) FindByUsername(ctx context.Context, username string, user *Model) error {
	return repo.db.QueryRow(ctx,
		`
			SELECT 
				id, username, password
			FROM
				user
			WHERE username = $1 
		`,
		username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash)
}
