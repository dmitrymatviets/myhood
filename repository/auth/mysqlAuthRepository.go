package auth

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
)

type MysqlAuthRepository struct {
	db database.Database
}

func (MysqlAuthRepository) CheckCredentials(ctx context.Context, credentials contract.Credentials) (bool, error) {
	panic("implement me")
}

func (MysqlAuthRepository) StartSession(ctx context.Context, userId model.IntId) (contract.Session, error) {
	panic("implement me")
}

func (MysqlAuthRepository) GetUserIdBySession(ctx context.Context, sessionId contract.Session) (model.IntId, error) {
	panic("implement me")
}

func (MysqlAuthRepository) Logout(ctx context.Context, sessionId contract.Session) error {
	panic("implement me")
}
