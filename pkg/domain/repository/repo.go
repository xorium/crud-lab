package repository

import (
	"context"
	"crud-lab/pkg/domain/entities"
	"crud-lab/pkg/filters"
)

type UsersRepo interface {
	List(ctx context.Context, filter filters.Filter) ([]*entities.User, error)
	Get(ctx context.Context, id string) (*entities.User, error)
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
}

type Repository interface {
	Users() UsersRepo
}
