package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
)

type IUserService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto model.SignupDto) (*model.User, error)
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение списка друзей
	GetFriends(ctx context.Context, userId model.IntId) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (model.IntId, error)
}

type IUserRepository interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto model.SignupDto) (*model.User, error)
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение списка друзей
	GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error)
	// сохранение пользователя
	SaveUser(ctx context.Context, user *model.User) (model.IntId, error)
}
