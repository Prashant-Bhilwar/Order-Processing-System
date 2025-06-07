package repository

import (
	"database/sql"
	"user-service/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email =$1`
	row := r.DB.QueryRow(query, email)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}
