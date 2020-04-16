package test

import (
	"context"
	"fmt"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/repository/city"
	"github.com/dmitrymatviets/myhood/repository/user"
	"github.com/dmitrymatviets/myhood/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	rand2 "math/rand"
	"testing"
	"time"
)

func getAuthService() contract.IAuthService {
	var authService contract.IAuthService
	fx.New(
		fx.Provide(
			config.Load,
			database.NewDatabase,
			logger.New,
			city.NewMssqlCityRepository,
			user.NewMssqlUserRepository,
			service.NewAuthService,
		),
		fx.Populate(&authService),
	)
	return authService
}

func getValidSignupDto() model.SignupDto {
	rand2.Seed(time.Now().UnixNano())
	return model.SignupDto{
		Credentials: model.Credentials{
			Email:    fmt.Sprintf("test%d@test%d.com", rand2.Int(), rand2.Int()),
			Password: "12345",
		},
		Name:        "Дмитрий",
		Surname:     "Матвиец",
		DateOfBirth: time.Date(1989, 07, 17, 00, 00, 00, 00, time.Local),
		Gender:      "м",
		Interests:   []string{"программирование"},
		CityId:      1,
	}
}

func TestSignup_ValidSignupDto_CreatesUser(t *testing.T) {
	as := getAuthService()
	session, user, err := as.SignUp(context.Background(), getValidSignupDto())
	assert.Nil(t, err)
	assert.NotEmpty(t, session)
	assert.NotNil(t, user)
}

func TestSignup_DuplicatedEmail_Fails(t *testing.T) {
	as := getAuthService()
	dto1 := getValidSignupDto()
	session, user, err := as.SignUp(context.Background(), dto1)
	assert.Nil(t, err)
	assert.NotEmpty(t, session)
	assert.NotNil(t, user)
	dto2 := getValidSignupDto()
	dto2.Credentials.Email = dto1.Credentials.Email
	session, user, err = as.SignUp(context.Background(), dto2)
	assert.NotNil(t, err)
	assert.Empty(t, session)
	assert.Nil(t, user)
}

func TestSignup_NoEmail_Fails(t *testing.T) {
	as := getAuthService()
	dto1 := getValidSignupDto()
	dto1.Email = ""
	session, user, err := as.SignUp(context.Background(), dto1)
	assert.NotNil(t, err)
	assert.Empty(t, session)
	assert.Nil(t, user)
}

func TestSignup_NoPass_Fails(t *testing.T) {
	as := getAuthService()
	dto1 := getValidSignupDto()
	dto1.Password = ""
	session, user, err := as.SignUp(context.Background(), dto1)
	assert.NotNil(t, err)
	assert.Empty(t, session)
	assert.Nil(t, user)
}
