package database

import (
	"EM-TASK/app/internal/domain"
	"context"
)

type Database interface {
	CreateUser(ctx context.Context, req *domain.UserResponse) (*domain.UserResponse, error)
	GetUser(ctx context.Context, req *domain.GetUserByIdRequest) (*domain.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]domain.UserResponse, error)
	GetAllSortUsers(ctx context.Context, req *domain.UserSort) ([]domain.UserResponse, error)
	UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) error
	DeleteUser(ctx context.Context, req *domain.GetUserByIdRequest) error
	//Close()
}
