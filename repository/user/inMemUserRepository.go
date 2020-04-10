package auth

import (
	"context"
	"github.com/dmitrymatviets/myhood/pkg"
	"github.com/pkg/errors"
	"sync"

	"github.com/dmitrymatviets/myhood/core/model"
)

type InmemoryUserRepository struct {
	sessions    map[model.Session]model.IntId
	credentials map[model.Credentials]model.IntId
	sync.RWMutex
}

func (r *InmemoryUserRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func (r *InmemoryUserRepository) Authenticate(ctx context.Context, credentials model.Credentials) (model.IntId, error) {
	r.RLock()
	defer r.RUnlock()

	if id, ok := r.credentials[credentials]; ok {
		return id, nil
	}
	return 0, errors.New("invalid credentials")
}

func (r *InmemoryUserRepository) GetUserIdBySession(ctx context.Context, sessionId model.Session) (model.IntId, error) {
	r.RLock()
	defer r.RUnlock()

	if id, ok := r.sessions[sessionId]; ok {
		return id, nil
	}
	return 0, errors.New("invalid sessionId")
}

func (r *InmemoryUserRepository) Logout(ctx context.Context, sessionId model.Session) error {
	r.Lock()
	defer r.Unlock()

	if _, err := r.GetUserIdBySession(ctx, sessionId); err != nil {
		delete(r.sessions, sessionId)
		return nil
	}
	return errors.New("session does not exist")
}

func (r *InmemoryUserRepository) SignUp(ctx context.Context, dto model.SignupDto) (model.Session, *model.User, *pkg.PublicError) {
	var session model.Session
	var user *model.User

	var err error
	user, err = as.userRepo.SignUp(ctx, dto)
	if err != nil {
		return pkg.NewPublicError("Ошибка регистрации", err)
	}
	session, err = as.authRepo.StartSession(ctx, user)
	if err != nil {
		return pkg.NewPublicError("Ошибка входа", err)
	}
	// TODO нотификация
	return nil

	if err != nil {
		return session, user, err
	}

	return session, user, err
	return
}

func (r *InmemoryUserRepository) GetById(ctx context.Context, id model.IntId) (*model.User, error) {
	panic("implement me")
}

func (r *InmemoryUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("implement me")
}

func (r *InmemoryUserRepository) GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error) {
	panic("implement me")
}

func (r *InmemoryUserRepository) GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error) {
	panic("implement me")
}

func (r *InmemoryUserRepository) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("implement me")
}

func (r *InmemoryUserRepository) startSession(ctx context.Context, user *model.User) (model.Session, error) {
	r.Lock()
	defer r.Unlock()

	session := model.NewSession()
	r.sessions[session] = user.Id
	return session, nil
}
