package user

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
)

type MysqlUserRepository struct{}

func (MysqlUserRepository) SignUp(ctx context.Context, dto contract.SignupDto) (*model.User, error) {
	panic("implement me")
}

func (MysqlUserRepository) GetById(ctx context.Context, id model.IntId) (*model.User, error) {
	panic("implement me")
}

func (MysqlUserRepository) GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error) {
	panic("implement me")
}

func (MysqlUserRepository) GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error) {
	panic("implement me")
}

func (MysqlUserRepository) SaveUser(ctx context.Context, user *model.User) (model.IntId, error) {
	panic("implement me")
}
