package repository

import (
	"belajargolang/entity"
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	query := "INSERT INTO users(username, password) VALUES(?, ?)"
	_, err := repository.DB.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, id int32) (entity.User, error) {
	user := entity.User{}

	query := "SELECT id, username, password FROM users WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if !rows.Next() {
		return user, errors.New("ID NOT FOUND")
	}

	rows.Scan(&user.Id, &user.Username, &user.Password)
	return user, nil
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	query := "SELECT id, username, password FROM users"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	return users, nil
}
