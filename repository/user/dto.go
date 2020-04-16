package user

import (
	"encoding/json"
	"github.com/dmitrymatviets/myhood/core/model"
	"time"
)

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
}

func newUserDtoFromUser(u *model.User) userDto {
	interestsJson, err := json.Marshal(u.Interests)
	if err != nil || u.Interests == nil {
		interestsJson = []byte("[]")
	}

	return userDto{
		Id:            u.Id,
		Email:         u.Email,
		Name:          u.Name,
		Surname:       u.Surname,
		DateOfBirth:   u.DateOfBirth,
		Gender:        u.Gender,
		Interests:     string(interestsJson),
		CityId:        u.CityId,
		PageSlug:      u.Page.Slug,
		PageIsPrivate: u.Page.IsPrivate,
	}
}

func (u *userDto) toUser() *model.User {
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
