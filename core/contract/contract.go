package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/pkg"
)

type IAuthService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, *pkg.PublicError)
	// аутентификация пользователя
	Login(ctx context.Context, credentials model.Credentials) (model.Session, *pkg.PublicError)
	// получение пользователя по сессии
	GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, *pkg.PublicError)
	// выход из системы
	Logout(ctx context.Context, sessionId model.Session) *pkg.PublicError
}

type IUserService interface {
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, *pkg.PublicError)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, *pkg.PublicError)
	// получение списка друзей
	GetFriends(ctx context.Context, userId model.IntId) ([]*model.DisplayUserDto, *pkg.PublicError)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (*model.User, *pkg.PublicError)
}

type IUserRepository interface {
	// регистрация пользователя
	SignUp(ctx context.Context, UserWithPassword *model.UserWithPassword) (model.Session, *model.User, error)
	// аутентификация
	Authenticate(ctx context.Context, credentials model.Credentials) (model.Session, *model.User, error)
	// получение id пользователя по сессии
	GetUserIdBySession(ctx context.Context, sessionId model.Session) (model.IntId, error)
	// выход из системы
	Logout(ctx context.Context, sessionId model.Session) error
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение пользователя по email
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение списка друзей
	GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
}

type ICityRepository interface {
	// получение списка городов
	GetCities(ctx context.Context) ([]*model.City, error)
	// получение города по id
	GetById(ctx context.Context, id model.IntId) (*model.City, error)
}

/*
type IMessageService interface {
	SendMessage(ctx context.Context, from model.User, to model.IntId, msg string) error
	MarkMessageAsRead(ctx context.Context, msgId string) error
	GetThreads(ctx context.Context, user *model.User, count int) ([]*model.Thread, error)
	GetMessages(ctx context.Context, user *model.User, threadId model.IntId, count int) ([]*model.Message, error)
}

type IMessageRepository interface {
	SendMessage(ctx context.Context, msg model.Message, from model.User, to model.User) error
	MarkMessageAsRead(ctx context.Context, msg model.Message) error
	GetThreads(ctx context.Context, user *model.User, count int) ([]*model.Thread, error)
	GetMessages(ctx context.Context, thread *model.Thread, count int) ([]*model.Message, error)
}
*/

/*
type ITransactional interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) (err error)
}
*/
