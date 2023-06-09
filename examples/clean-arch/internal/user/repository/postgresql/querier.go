// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"github.com/aliml92/sqlc-customizer/examples/clean-arch/internal/models"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

var _ Querier = (*Queries)(nil)
