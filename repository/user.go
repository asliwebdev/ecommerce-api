package repository

import (
	"database/sql"
	"fmt"

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

	tx, err := u.DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Commit()

	_, err = tx.Exec(`INSERT INTO users (id, name, role, email)
	 VALUES ($1, $2, $3, $4)`,
		id, user.Name, user.Role, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.QueryRow(`SELECT id, name, role, email FROM users WHERE email = $1`, email).Scan(&user.Id, &user.Name, &user.Role, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetUserById(userId string) (*models.User, error) {
	user := &models.User{}
	err := u.DB.QueryRow(`SELECT id, name, role, email, password, created_at, updated_at FROM users WHERE id = $1`, userId).Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error querying user by ID: %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	rows, err := u.DB.Query(`SELECT id, name, role, email, password, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("error querying all users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) UserExists(userId string) (bool, error) {
	var exists bool
	err := u.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, userId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}
	return exists, nil
}

func (u *UserRepo) UpdateUser(id string, user models.User) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Commit()

	_, err = tx.Exec(`UPDATE users SET name = $1, role = $2, email = $3, password = $4 WHERE id = $5`,
		user.Name, user.Role, user.Email, user.Password, id)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (u *UserRepo) DeleteUser(userID string) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM users WHERE id = $1`, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return tx.Commit()
}
