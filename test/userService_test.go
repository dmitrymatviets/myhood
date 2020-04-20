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
	"math/rand"
	"sync"
	"testing"
	"time"
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

//region getByIds
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

//endregion

//region friends
func TestGetFriends_BadSession_Fails(t *testing.T) {
	us := getUserService()
	friends, err := us.GetFriends(context.Background(), "badSession", 1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "сессия")
	assert.Len(t, friends, 0)
	fmt.Println(err)
}

func TestGetFriends_Signup_ZeroFriends(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	friends, err := us.GetFriends(context.Background(), session, user.Id)
	assert.Nil(t, err)
	assert.Len(t, friends, 0)
}

func TestAddFriend_BadSession_Fails(t *testing.T) {
	us := getUserService()
	err := us.AddFriend(context.Background(), "badSession", 1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "сессия")
	fmt.Println(err)
}

func TestAddFriend_InvalidId_Fails(t *testing.T) {
	us := getUserService()
	session, _ := createValidUser()
	err := us.AddFriend(context.Background(), session, -1)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestAddFriend_SelfAdd_Fails(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	err := us.AddFriend(context.Background(), session, user.Id)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestAddFriend_CorrectId_Success(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	_, friend := createValidUser()
	err := us.AddFriend(context.Background(), session, friend.Id)
	assert.NoError(t, err)
	friends, err := us.GetFriends(context.Background(), session, user.Id)
	assert.NoError(t, err)
	assert.Len(t, friends, 1)
	assert.Equal(t, friends[0].Id, friend.Id)
	friendsOfFriend, err := us.GetFriends(context.Background(), session, friend.Id)
	assert.NoError(t, err)
	assert.Empty(t, friendsOfFriend)
}

func TestRemoveFriend_BadSession_Fails(t *testing.T) {
	us := getUserService()
	err := us.RemoveFriend(context.Background(), "badSession", 1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "сессия")
	fmt.Println(err)
}

func TestRemoveFriend_BadId_Fail(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	err := us.RemoveFriend(context.Background(), session, -1)
	assert.Error(t, err)
	friendsOfFriend, err := us.GetFriends(context.Background(), session, user.Id)
	assert.NoError(t, err)
	assert.Empty(t, friendsOfFriend)
}

func TestRemoveFriend_CorrectId_Success(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	_, friend := createValidUser()
	err := us.AddFriend(context.Background(), session, friend.Id)
	assert.NoError(t, err)
	friends, err := us.GetFriends(context.Background(), session, user.Id)
	assert.NoError(t, err)
	assert.Len(t, friends, 1)
	err = us.RemoveFriend(context.Background(), session, friend.Id)
	assert.NoError(t, err)
	friends, err = us.GetFriends(context.Background(), session, user.Id)
	assert.NoError(t, err)
	assert.Empty(t, friends)
}

//endregion

//region saveUser
func TestSaveUser_BadSession_Fail(t *testing.T) {
	us := getUserService()
	_, user := createValidUser()
	user, err := us.SaveUser(context.Background(), "badSession", user)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestSaveUser_BadId_Fail(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()
	user.Id = -1
	user, err := us.SaveUser(context.Background(), session, user)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestSaveUser_CorrectId_Success(t *testing.T) {
	us := getUserService()
	session, user := createValidUser()

	rand.Seed(time.Now().UnixNano())

	user.Email = user.Email + ".test"
	user.Name = "Иван"
	user.Surname = "Иванов"
	user.DateOfBirth = time.Date(1991, 12, 12, 11, 11, 11, 00, time.Local)
	user.Interests = []string{"пение", "гитара"}
	user.CityId = 2
	user.Page.Slug = fmt.Sprintf("page-%d", rand.Int())
	user.Page.IsPrivate = true
	savedUser, err := us.SaveUser(context.Background(), session, user)
	assert.NoError(t, err)
	assert.NotNil(t, savedUser)
	assert.Equal(t, savedUser.Id, user.Id)
	assert.Equal(t, savedUser.Email, user.Email)
	assert.Equal(t, savedUser.Name, user.Name)
	assert.Equal(t, savedUser.Surname, user.Surname)
	assert.Equal(t, savedUser.DateOfBirth.UnixNano(), user.DateOfBirth.UnixNano())
	assert.Equal(t, savedUser.Interests, user.Interests)
	assert.Equal(t, savedUser.CityId, user.CityId)
	assert.Equal(t, savedUser.Page, user.Page)
}

//endregion

func TestGetRecommendations_BadSession_Fail(t *testing.T) {
	us := getUserService()
	recs, err := us.GetRecommendations(context.Background(), "badSession")
	assert.Error(t, err)
	assert.Empty(t, recs)
}

func TestGetRecommendations_CorrectSession_Success(t *testing.T) {
	us := getUserService()
	session, _ := createValidUser()
	recs, err := us.GetRecommendations(context.Background(), session)
	assert.NoError(t, err)
	assert.NotEmpty(t, recs)
}
