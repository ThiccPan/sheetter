package usecase

import (
	"github.com/thiccpan/sheetter/internal/entity"
	"github.com/thiccpan/sheetter/internal/usecase/webapi"
)

type userUsecase struct {
	sheetApi webapi.SheetApi
}

func NewUserUsecase(sa webapi.SheetApi) UserUsecase {
	return &userUsecase{
		sheetApi: sa,
	}
}

func (uu *userUsecase) GetAllData() ([]entity.User, error) {
	data, err := uu.sheetApi.ReadFromSheet()
	return data, err
}