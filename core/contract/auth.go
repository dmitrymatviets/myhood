package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
)

type IAuthService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, error)
	// аутентификация пользователя
	Login(ctx context.Context, credentials model.Credentials) (model.Session, error)
	// получение пользователя по сессии
	GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, error)
	// выход из системы
	Logout(ctx context.Context, sessionId model.Session) error
}

type IAuthRepository interface {
	// аутентификация
	Authenticate(ctx context.Context, credentials model.Credentials) (model.IntId, error)
	// начало сессии
	StartSession(ctx context.Context, userId *model.User) (model.Session, error)
	// получение id пользователя по сессии
	GetUserIdBySession(ctx context.Context, sessionId model.Session) (model.IntId, error)
	// выход из системы
	Logout(ctx context.Context, sessionId model.Session) error
}
