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
	var session model.Session
	var user *model.User

	err := as.validateSignupDto(ctx, dto)
	if err != nil {
		return session, user, err
	}

	city, err := as.cityRepo.GetById(ctx, dto.CityId)
	if city == nil {
		return session, user, pkg.NewPublicError("Неверный город")
	}
	if err != nil {
		return session, user, pkg.NewPublicError("Ошибка проверки города", err)
	}

	session, user, err:=as.userRepo.SignUp(ctx, dto){

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

func (as *AuthService) validateSignupDto(ctx context.Context, dto model.SignupDto) error {
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
