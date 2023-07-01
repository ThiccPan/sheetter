package usecase

import (
	"context"

	"github.com/thiccpan/sheetter/internal/entity"
)

type UserUsecase interface {
	GetAllData(ctx context.Context) ([]entity.User, error)
	CreateData(ctx context.Context, data entity.User) (entity.User, error)
}
