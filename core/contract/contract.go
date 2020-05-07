package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
	"time"
)

type IAuthService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, error)
	// аутентификация пользователя
	Login(ctx context.Context, credentials model.Credentials) (model.Session, *model.User, error)
	// получение пользователя по сессии.
	// возвращает ошибку в любом неуспешном случае.
	GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, error)
	// выход из системы
	Logout(ctx context.Context, sessionId model.Session) error
	// очистка устаревших сессий
	CleanSessions(ctx context.Context, lifeDurationThreshold time.Duration)
}

type IUserService interface {
	// получение пользователя по id
	GetById(ctx context.Context, sessionId model.Session, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, sessionId model.Session, ids []model.IntId) ([]*model.User, error)
	// поиск пользователей
	Search(ctx context.Context, sessionId model.Session, searchDto model.SearchDto) ([]*model.DisplayUserDto, error)
	// получение списка друзей
	GetFriends(ctx context.Context, sessionId model.Session, userId model.IntId) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, sessionId model.Session, user *model.User) (*model.User, error)
	// добавление друга
	AddFriend(ctx context.Context, sessionId model.Session, friendId model.IntId) error
	// удаление друга
	RemoveFriend(ctx context.Context, sessionId model.Session, friendId model.IntId) error
	// получение рекомендаций для добавления в друзья
	GetRecommendations(ctx context.Context, sessionId model.Session) ([]*model.DisplayUserDto, error)
}

type ICityService interface {
	// получение списка городов
	GetCities(ctx context.Context) ([]*model.City, error)
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
	// очистка устаревших сессий
	CleanSessions(ctx context.Context, lifeDurationThreshold time.Duration)
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение пользователя по email
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// поиск пользователей
	Search(ctx context.Context, searchDto model.SearchDto) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
	// список друзей
	GetFriends(ctx context.Context, userId *model.User) ([]*model.DisplayUserDto, error)
	// добавление друга
	AddFriend(ctx context.Context, user *model.User, friend *model.User) error
	// удаление друга
	RemoveFriend(ctx context.Context, user *model.User, friend *model.User) error
	// получение рекомендаций для добавления в друзья
	GetRecommendations(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error)
}

type ICityRepository interface {
	// получение списка городов
	GetCities(ctx context.Context) ([]*model.City, error)
	// получение города по id
	GetById(ctx context.Context, id model.IntId) (*model.City, error)
}
