package contract

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/model"
)

type IFriendService interface {
	// добавление в друзья
	AddFriend(ctx context.Context, user *model.User, friendUserId model.IntId) error
	// удаление из друзей
	RemoveFriend(ctx context.Context, user *model.User, friendUserId model.IntId) error
	// получение друзей
	GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUser, error)
}

type IFriendRepository interface {
	// добавление в друзья
	AddFriend(ctx context.Context, user *model.User, friend *model.User) error
	// удаление из друзей
	RemoveFriend(ctx context.Context, user *model.User, friend *model.User) error
	// получение друзей
	GetFriendIds(ctx context.Context, user *model.User) ([]model.IntId, error)
}
