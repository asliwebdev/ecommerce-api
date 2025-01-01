package repository

import (
	"database/sql"

	"ecommerce/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user models.User) error {
	id := uuid.NewString()
	_, err := u.DB.Exec(`INSERT INTO users (id, name, role, email)
	 VALUES ($1, $2, $3, $4)`,
		id, user.Name, user.Role, user.Email)
	if err != nil {
		return err
	}

	return nil
}
