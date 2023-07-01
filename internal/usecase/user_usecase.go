package usecase

import (
	"context"

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

func (uu *userUsecase) GetAllData(ctx context.Context) ([]entity.User, error) {
	data, err := uu.sheetApi.ReadFromSheet()
	return data, err
}

func (uu *userUsecase) CreateData(ctx context.Context, data entity.User) (entity.User, error) {
	err := uu.sheetApi.WriteToRow(data)
	if err != nil {
		return entity.User{}, err
	}
	return data, nil	
}