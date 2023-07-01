package usecase

import "github.com/thiccpan/sheetter/internal/entity"

type UserUsecase interface {
	GetAllData() ([]entity.User, error)
}