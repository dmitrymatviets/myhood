package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/validator"
	"github.com/dmitrymatviets/myhood/pkg"
)

type UserService struct {
	userRepo    contract.IUserRepository
	authService contract.IAuthService
	validator   *validator.Validator
}

func NewUserService(userRepo contract.IUserRepository, authService contract.IAuthService, validator *validator.Validator) contract.IUserService {
	return &UserService{userRepo: userRepo, authService: authService, validator: validator}
}

func (us *UserService) GetById(ctx context.Context, sessionId model.Session, id model.IntId) (*model.User, error) {
	_, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return us.userRepo.GetById(ctx, id)
}

func (us *UserService) GetByIds(ctx context.Context, sessionId model.Session, ids []model.IntId) ([]*model.User, error) {
	_, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return us.userRepo.GetByIds(ctx, ids)
}

func (us *UserService) GetFriends(ctx context.Context, sessionId model.Session, userId model.IntId) ([]*model.DisplayUserDto, error) {
	user, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	if userId != user.Id {
		user, err = us.userRepo.GetById(ctx, userId)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, pkg.NewPublicError("Ошибка получения пользователя")
	}

	return us.userRepo.GetFriends(ctx, user)
}

func (us *UserService) SaveUser(ctx context.Context, sessionId model.Session, user *model.User) (*model.User, error) {
	sessionUser, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	if sessionUser.Id != user.Id {
		return nil, pkg.NewPublicError("Некорректный пользователь")
	}

	err = us.validator.ValidateStruct(user)
	if err != nil {
		return nil, err
	}

	return us.userRepo.SaveUser(ctx, user)
}

func (us *UserService) AddFriend(ctx context.Context, sessionId model.Session, friendId model.IntId) error {
	sessionUser, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return err
	}

	if friendId == sessionUser.Id {
		return pkg.NewPublicError("Нельзя добавить в друзья самого себя")
	}

	friend, err := us.userRepo.GetById(ctx, friendId)
	if err != nil {
		return err
	}

	if friend == nil {
		return pkg.NewPublicError("Ошибка получения друга")
	}

	existingFriends, err := us.userRepo.GetFriends(ctx, sessionUser)
	if err != nil {
		return err
	}

	for _, existingFriend := range existingFriends {
		if existingFriend.Id == friendId {
			return pkg.NewPublicError("Данный человек уже добавлен в друзья")
		}
	}

	return us.userRepo.AddFriend(ctx, sessionUser, friend)
}

func (us *UserService) RemoveFriend(ctx context.Context, sessionId model.Session, friendId model.IntId) error {
	sessionUser, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return err
	}

	friend, err := us.userRepo.GetById(ctx, friendId)
	if err != nil {
		return err
	}

	if friend == nil {
		return pkg.NewPublicError("Ошибка получения друга")
	}

	return us.userRepo.RemoveFriend(ctx, sessionUser, friend)
}

func (us *UserService) GetRecommendations(ctx context.Context, sessionId model.Session) ([]*model.DisplayUserDto, error) {
	sessionUser, err := us.authService.GetUserBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	return us.userRepo.GetRecommendations(ctx, sessionUser)
}
