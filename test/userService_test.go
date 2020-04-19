package test

import (
	"context"
	"fmt"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/validator"
	"github.com/dmitrymatviets/myhood/repository/city"
	"github.com/dmitrymatviets/myhood/repository/user"
	"github.com/dmitrymatviets/myhood/service"
	assert "github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"sync"
	"testing"
)

var userService contract.IUserService
var onceUser sync.Once

func getUserService() contract.IUserService {
	onceUser.Do(func() {
		fx.New(
			fx.NopLogger,
			fx.Provide(
				config.Load,
				database.NewDatabase,
				logger.New,
				validator.NewValidator,
				city.NewMssqlCityRepository,
				user.NewMssqlUserRepository,
				service.NewAuthService,
				service.NewUserService,
			),
			fx.Populate(&userService),
		)
	})
	return userService
}

//region getById
func TestGetById_BadSession_Fails(t *testing.T) {
	us := getUserService()
	user, err := us.GetById(context.Background(), "badSession", 1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "сессия")
	assert.Nil(t, user)
	fmt.Println(err)
}

func TestGetById_CorrectUser_Success(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	gotUser, err := us.GetById(context.Background(), session, user.Id)
	assert.Nil(t, err)
	assert.NotNil(t, gotUser)
	assert.Equal(t, user.Id, gotUser.Id)
}

func TestGetById_BadUser_Fail(t *testing.T) {
	us := getUserService()
	session, _ := createValidUser()
	gotUser, err := us.GetById(context.Background(), session, -1)
	assert.Nil(t, err)
	assert.Nil(t, gotUser)
}

//endregion

func TestGetByIds_BadSession_Fails(t *testing.T) {
	us := getUserService()
	user, err := us.GetByIds(context.Background(), "badSession", []model.IntId{1})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "сессия")
	assert.Nil(t, user)
	fmt.Println(err)
}

func TestGetByIds_WrongIds_Fails(t *testing.T) {
	us := getUserService()
	session, _ := createValidUser()
	users, err := us.GetByIds(context.Background(), session, []model.IntId{-1})
	assert.Nil(t, err)
	assert.Len(t, users, 0)
}

func TestGetByIds_CorrectIds_Success(t *testing.T) {
	us := getUserService()
	_, user1 := createValidUser()
	session, user2 := createValidUser()
	users, err := us.GetByIds(context.Background(), session, []model.IntId{user1.Id, user2.Id})
	assert.Nil(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, users[0].Id, user1.Id)
	assert.Equal(t, users[1].Id, user2.Id)
}
