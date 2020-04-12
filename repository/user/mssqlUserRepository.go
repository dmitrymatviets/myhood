package user

import (
	"context"
	"encoding/json"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/pkg"
	"github.com/pkg/errors"
	"time"
)

const hashSalt = "myhood_93284203948092384767709324890128309128!@#!@#!@$%%^^^"

type MssqlUserRepository struct {
	db *database.Database
}

func NewMssqlUserRepository(db *database.Database) contract.IUserRepository {
	return &MssqlUserRepository{db: db}
}

func (ur *MssqlUserRepository) SignUp(ctx context.Context, user *model.UserWithPassword) (model.Session, *model.User, error) {
	var sessionId model.Session
	var savedUser *model.User

	err := ur.db.WithTransaction(ctx, func(ctx context.Context) error {
		var localErr error
		savedUser, localErr = ur.createUser(ctx, user)
		if localErr != nil {
			return pkg.NewPublicError("Ошибка регистрации", localErr)
		}

		sessionId, savedUser, localErr = ur.Authenticate(ctx, model.Credentials{Email: user.Email, Password: user.Password})
		if localErr != nil {
			return pkg.NewPublicError("Ошибка завершения регистрации", localErr)
		}

		return nil
	})

	if err != nil {
		return "", nil, err
	}

	return sessionId, savedUser, nil
}

func (ur *MssqlUserRepository) Authenticate(ctx context.Context, credentials model.Credentials) (model.Session, *model.User, error) {
	var sessionId model.Session
	var user *model.User

	tErr := ur.db.WithTransaction(ctx, func(ctx context.Context) error {
		userId, err := ur.authenticateInternal(ctx, credentials)
		if err != nil {
			return pkg.NewPublicError("Ошибка аутентификации", err)
		}
		user, err = ur.GetById(ctx, userId)
		if err != nil {
			return pkg.NewPublicError("Ошибка завершения аутентификации", err)
		}
		sessionId, err = ur.startSession(ctx, userId)
		if err != nil {
			return pkg.NewPublicError("Ошибка начала сессии", err)
		}
		return nil
	})

	if tErr != nil {
		return "", nil, tErr
	}

	return sessionId, user, nil
}

func (ur *MssqlUserRepository) GetUserIdBySession(ctx context.Context, sessionId model.Session) (model.IntId, error) {
	panic("implement me")
}

func (ur *MssqlUserRepository) Logout(ctx context.Context, sessionId model.Session) error {
	panic("implement me")
}

func (ur *MssqlUserRepository) GetById(ctx context.Context, id model.IntId) (*model.User, error) {
	dtoUser := userDto{}
	err := ur.db.TxOrDbFromContext(ctx).GetContext(ctx, &dtoUser,
		`select user_id
                     , email
                     , name
                     , surname
                     , date_of_birth
                     , gender
                     , interests
                     , city_id
                     , page_slug
                     , page_is_private     
                  from users     
                 where user_id = ?`,
		id)

	if err != nil {
		return nil, pkg.NewPublicError("Ошибка получения пользователя", err)
	}

	return dtoUser.toUser(), nil
}

func (ur *MssqlUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("implement me")
}

func (ur *MssqlUserRepository) GetByIds(ctx context.Context, ids []model.IntId) ([]*model.User, error) {
	panic("implement me")
}

func (ur *MssqlUserRepository) GetFriends(ctx context.Context, user *model.User) ([]*model.DisplayUserDto, error) {
	panic("implement me")
}

func (ur *MssqlUserRepository) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("implement me")
}

func (ur *MssqlUserRepository) createUser(ctx context.Context, user *model.UserWithPassword) (*model.User, error) {

	interestsJson, _ := json.Marshal(user.Interests)

	result, err := ur.db.TxOrDbFromContext(ctx).ExecContext(ctx, `insert into users(email, hash, name, surname, date_of_birth, gender, interests, city_id, page_slug, page_is_private) 
                                                  values (?, md5(?), ?, ?, ?, ?, ?, ?, ? , ?)`,
		user.Email,
		hashSalt+user.Password,
		user.Name,
		user.Surname,
		user.DateOfBirth,
		user.Gender,
		interestsJson,
		user.CityId,
		user.Page.Slug,
		false,
	)

	if err != nil {
		return nil, pkg.NewPublicError("Ошибка создания пользователя", errors.WithStack(err))
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, pkg.NewPublicError("Ошибка завершения создания пользователя", errors.WithStack(err))
	}

	user.Id = model.IntId(lastInsertId)

	return user.User, nil
}

func (ur *MssqlUserRepository) authenticateInternal(ctx context.Context, credentials model.Credentials) (model.IntId, error) {
	var userId model.IntId
	err := ur.db.TxOrDbFromContext(ctx).GetContext(ctx, &userId,
		`select user_id from users
                where email = ? and hash = md5(?)`,
		credentials.Email,
		hashSalt+credentials.Password,
	)

	if err != nil {
		return 0, pkg.NewPublicError("Ошибка аутентификации", errors.WithStack(err))
	}

	return userId, nil
}

func (ur *MssqlUserRepository) startSession(ctx context.Context, userId model.IntId) (model.Session, error) {
	sessionId := model.NewSession()
	_, err := ur.db.TxOrDbFromContext(ctx).ExecContext(ctx,
		`insert into sessions(session_id, user_id, created) 
                values (?, ?, now())`,
		sessionId,
		userId,
	)

	if err != nil {
		return "", pkg.NewPublicError("Ошибка начала сессии", errors.WithStack(err))
	}

	return sessionId, nil
}

type userDto struct {
	Id            model.IntId `db:"user_id"`
	Email         string      `db:"email"`
	Name          string      `db:"name"`
	Surname       string      `db:"surname"`
	DateOfBirth   time.Time   `db:"date_of_birth"`
	Gender        string      `db:"gender"`
	Interests     string      `db:"interests"`
	CityId        model.IntId `db:"city_id"`
	PageSlug      string      `db:"page_slug"`
	PageIsPrivate bool        `db:"page_is_private"`
	//TODO
	FriendIds []model.IntId
}

func (u userDto) toUser() *model.User {
	interests := make([]string, 0)
	_ = json.Unmarshal([]byte(u.Interests), &interests)

	return &model.User{
		Id:          u.Id,
		Email:       u.Email,
		Name:        u.Name,
		Surname:     u.Surname,
		DateOfBirth: u.DateOfBirth,
		Gender:      u.Gender,
		Interests:   interests,
		CityId:      u.CityId,
		Page: model.Page{
			Slug:      u.PageSlug,
			IsPrivate: u.PageIsPrivate,
		},
	}
}