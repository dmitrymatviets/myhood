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

type IUserService interface {
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение списка друзей
	GetFriends(ctx context.Context, userId model.IntId) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
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
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
	// добавление друга
	AddFriend(ctx context.Context, user *model.User, friend *model.User) error
	// удаление друга
	RemoveFriend(ctx context.Context, user *model.User, friend *model.User) error
}

type ICityRepository interface {
	// получение списка городов
	GetCities(ctx context.Context) ([]*model.City, error)
	// получение города по id
	GetById(ctx context.Context, id model.IntId) (*model.City, error)
}
