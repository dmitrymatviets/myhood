package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
	"time"
)

type IUserService interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto SignupDto) (*model.User, error)
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение пользователя по email
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type IUserRepository interface {
	// регистрация пользователя
	SignUp(ctx context.Context, dto SignupDto) (*model.User, error)
	// получение пользователя по id
	GetById(ctx context.Context, id model.IntId) (*model.User, error)
	// получение нескольких пользователей по id
	GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error)
	// получение пользователя по email
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

// DTO для регистрации пользователя
type SignupDto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Interests   []string  `json:"interests"`
	City        string    `json:"city"`
}
