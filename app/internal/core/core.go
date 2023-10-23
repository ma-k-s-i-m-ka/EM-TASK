package core

import (
	"EM-TASK/app/internal/domain"
	"EM-TASK/app/internal/infrastructure/logger"
	"EM-TASK/app/internal/repository/database"
	"EM-TASK/app/internal/repository/database/postgresql"
	"context"
	"fmt"
	"os"
)

// Core represents business logic layer interface.
type Core interface {
	CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.UserResponse, error)
	GetUser(ctx context.Context, req *domain.GetUserByIdRequest) (*domain.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]domain.UserResponse, error)
	GetAllSortUsers(ctx context.Context, req *domain.UserSort) ([]domain.UserResponse, error)
	UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) (*domain.UserResponse, error)
	DeleteUser(ctx context.Context, req *domain.GetUserByIdRequest) error
}

// core implements Core interface.
type core struct {
	repo database.Database
}

// New returns Core instance.
func New(ctx context.Context) (Core, error) {
	db, err := postgresql.NewDriver(ctx, os.Getenv("DATABASE_DSN"))
	if err != nil {
		return nil, fmt.Errorf("creating postgresql driver: %w", err)
	}
	logger.Logger.Info().Msg("Connected to database")

	return &core{repo: db}, nil
}
