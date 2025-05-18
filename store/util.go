package store

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	var i User

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return i, fmt.Errorf("failed to hash password: %w", err)
	}

	row := q.db.QueryRow(ctx, createUserCommand, arg.Email, hashedPasswordBytes)

	err = row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invaild password")
	}
	return nil
}
