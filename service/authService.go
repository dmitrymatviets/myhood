package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	validator2 "github.com/dmitrymatviets/myhood/infrastructure/validator"
	"github.com/dmitrymatviets/myhood/pkg"
	"time"
)

type AuthService struct {
	cityRepo  contract.ICityRepository
	userRepo  contract.IUserRepository
	validator *validator2.Validator
}

func NewAuthService(cityRepo contract.ICityRepository, userRepo contract.IUserRepository, validator *validator2.Validator) contract.IAuthService {
	return &AuthService{
		cityRepo:  cityRepo,
		userRepo:  userRepo,
		validator: validator,
	}
}

func (as *AuthService) SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, error) {
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
	return as.userRepo.Authenticate(ctx, credentials)
}

func (as *AuthService) GetUserBySession(ctx context.Context, sessionId model.Session) (*model.User, error) {
	userId, err := as.userRepo.GetUserIdBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return nil, pkg.NewPublicError("Некорректная сессия")
	}
	return as.userRepo.GetById(ctx, userId)
}

func (as *AuthService) Logout(ctx context.Context, sessionId model.Session) error {
	return as.userRepo.Logout(ctx, sessionId)
}

func (as *AuthService) validateSignupDto(ctx context.Context, dto model.SignupDto) error {
	err := as.validator.ValidateStruct(dto)
	if err != nil {
		return err
	}

	if dto.DateOfBirth.After(time.Now().AddDate(-6, 0, 0)) {
		return pkg.NewValidationErr("регистрация возможна с 6 лет", err)
	}

	if dto.DateOfBirth.Before(time.Now().AddDate(-120, 0, 0)) {
		return pkg.NewValidationErr("некорректная дата", err)
	}

	city, err := as.cityRepo.GetById(ctx, dto.CityId)
	if err != nil {
		return pkg.NewPublicError("Ошибка проверки города", err)
	}
	if city == nil {
		return pkg.NewValidationErr("неверный город", nil)
	}
	return nil
}
