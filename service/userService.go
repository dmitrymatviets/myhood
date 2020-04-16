package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
)

type UserService struct {
	userRepo contract.IUserRepository
}

func NewUserService(userRepo contract.IUserRepository) contract.IUserService {
	return &UserService{userRepo: userRepo}
}

func (UserService) GetById(ctx context.Context, id model.IntId) (*model.User, error) {
	panic("implement me")
}

func (UserService) GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error) {
	panic("implement me")
}

func (UserService) GetFriends(ctx context.Context, userId model.IntId) ([]*model.DisplayUserDto, error) {
	panic("implement me")
}

func (UserService) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("implement me")
}
