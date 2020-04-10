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

	// валидация
	// https://github.com/gin-gonic/gin/issues/2167

	city, err := as.cityRepo.GetById()

	var session model.Session
	var user *model.User

	err := as.authRepo.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		user, err = as.userRepo.SignUp(ctx, dto)
		if err != nil {
			return pkg.NewPublicError("Ошибка регистрации", err)
		}
		session, err = as.authRepo.StartSession(ctx, user)
		if err != nil {
			return pkg.NewPublicError("Ошибка входа", err)
		}
		// TODO нотификация
		return nil
	})

	if err != nil {
		return session, user, err
	}

	return session, user, err
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

func validateSignupDto() {

}
