package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/pkg"
)

type AuthService struct {
	cityRepo contract.ICityRepository
	userRepo contract.IUserRepository
}

func (as *AuthService) SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, error) {
	err := as.validateSignupDto(ctx, dto)
	if err != nil {
		return "", nil, err
	}

	//дублирование емейлов
	session, user, err := as.userRepo.SignUp(ctx, dto)
	if err != nil {
		return "", nil, err
	}
}

func (as *AuthService) Login(ctx context.Context, credentials model.Credentials) (model.Session, error) {
	panic("implement me")
}

func (as *AuthService) GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, error) {
	panic("implement me")
}

func (as *AuthService) Logout(ctx context.Context, sessionId model.Session) error {
	panic("implement me")
}

func (as *AuthService) validateSignupDto(ctx context.Context, dto model.SignupDto) *pkg.PublicError {
	// TODO
	// валидация
	// https://github.com/gin-gonic/gin/issues/2167

	city, err := as.cityRepo.GetById(ctx, dto.CityId)
	if city == nil {
		return pkg.NewPublicError("Неверный город")
	}
	if err != nil {
		return pkg.NewPublicError("Ошибка проверки города", err)
	}
	return nil
}
