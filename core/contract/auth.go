package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
)

type IAuthService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto SignupDto) (Session, *model.User, error)
	// аутентификация пользователя
	Login(ctx context.Context, credentials Credentials) (Session, error)
	// получение пользователя по сессии
	GetUserBySession(ctx context.Context, sessionId Session) (*model.User, error)
	// выход из системы
	Logout(ctx context.Context, sessionId Session) error
}

type IAuthRepository interface {
	// начало сессии
	CheckCredentials(ctx context.Context, credentials Credentials) (bool, error)
	// начало сессии
	StartSession(ctx context.Context, userId model.IntId) (Session, error)
	// получение id пользователя по сессии
	GetUserIdBySession(ctx context.Context, sessionId Session) (model.IntId, error)
	// выход из системы
	Logout(ctx context.Context, sessionId Session) error
}

// DTO для аутентификации
type Credentials struct {
	Email    string
	Password string
}

// идентификатор сессии
type Session string
