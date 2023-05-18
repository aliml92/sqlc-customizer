// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"github.com/aliml92/searchfeed/internal/models"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    password 
) VALUES (
    $1,
    $2,
    $3,
    $4
) 
RETURNING id, username, email, password, bio, image, created_at, updated_at
`

type CreateUserParams struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*models.User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, bio, image, created_at, updated_at
FROM users 
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}