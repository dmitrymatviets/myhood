package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"

	"github.com/dmitrymatviets/myhood/core/model"
)

type InmemoryAuthRepository struct {
	sessions    map[model.Session]model.IntId
	credentials map[model.Credentials]model.IntId
	sync.RWMutex
}

func (r *InmemoryAuthRepository) CheckCredentials(ctx context.Context, credentials model.Credentials) (model.IntId, error) {
	r.RLock()
	defer r.RUnlock()

	if id, ok := r.credentials[credentials]; ok {
		return id, nil
	}
	return 0, errors.New("invalid credentials")
}

func (r *InmemoryAuthRepository) StartSession(ctx context.Context, user *model.User) (model.Session, error) {
	r.Lock()
	defer r.Unlock()

	session := model.Session(uuid.New().String())
	r.sessions[session] = user.Id
	return session, nil
}

func (r *InmemoryAuthRepository) GetUserIdBySession(ctx context.Context, sessionId model.Session) (model.IntId, error) {
	r.RLock()
	defer r.RUnlock()

	if id, ok := r.sessions[sessionId]; ok {
		return id, nil
	}
	return 0, errors.New("invalid sessionId")
}

func (r *InmemoryAuthRepository) Logout(ctx context.Context, sessionId model.Session) error {
	r.Lock()
	defer r.Unlock()

	if _, err := r.GetUserIdBySession(ctx, sessionId); err != nil {
		delete(r.sessions, sessionId)
		return nil
	}
	return errors.New("session does not exist")
}
