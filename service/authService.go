package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/pkg"
	"github.com/go-playground/validator/v10"
	"time"
)

type AuthService struct {
	cityRepo contract.ICityRepository
	userRepo contract.IUserRepository
}

func NewAuthService(cityRepo contract.ICityRepository, userRepo contract.IUserRepository) contract.IAuthService {
	return &AuthService{
		cityRepo: cityRepo,
		userRepo: userRepo,
	}
}

func (as *AuthService) SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, error) {
	//validation
	err := as.validateSignupDto(ctx, dto)
	if err != nil {
		return "", nil, err
	}

	existingUserForEmail, err := as.userRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", nil, err
	}

	if existingUserForEmail != nil {
		return "", nil, pkg.NewPublicError("Уже зарегистрирован пользователь с данным email")
	}

	return as.userRepo.SignUp(ctx, dto.ToUserWithPassword())
}

func (as *AuthService) Login(ctx context.Context, credentials model.Credentials) (model.Session, *model.User, error) {
	panic("implement me")
}

func (as *AuthService) GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, error) {
	panic("implement me")
}

func (as *AuthService) Logout(ctx context.Context, sessionId model.Session) error {
	panic("implement me")
}

func (as *AuthService) validateSignupDto(ctx context.Context, dto model.SignupDto) error {
	err := validator.New().Struct(dto)
	if err != nil {
		return pkg.NewPublicError("Ошибка валидации "+err.Error(), err)
	}

	if dto.DateOfBirth.After(time.Now().AddDate(-6, 0, 0)) {
		return pkg.NewPublicError("Ошибка валидации: регистрация возможна с 6 лет", err)
	}

	if dto.DateOfBirth.Before(time.Now().AddDate(-120, 0, 0)) {
		return pkg.NewPublicError("Ошибка валидации: некорректная дата", err)
	}

	city, err := as.cityRepo.GetById(ctx, dto.CityId)
	if city == nil {
		return pkg.NewPublicError("Ошибка валидации: неверный город")
	}
	if err != nil {
		return pkg.NewPublicError("Ошибка проверки города", err)
	}
	return nil
}
