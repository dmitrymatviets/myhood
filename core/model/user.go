package model

import (
	"github.com/google/uuid"
	"time"
)

type IntId int64

type User struct {
	Id          IntId     `json:"id"`
	Email       string    `json:"email" validate:"required,max=255,email"`
	Name        string    `json:"name" validate:"required,max=50"`
	Surname     string    `json:"surname" validate:"required,max=50"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"  validate:"required,oneof=м ж"`
	Interests   []string  `json:"interests"`
	CityId      IntId     `json:"cityId" validate:"required"`
	Page        Page      `json:"page"`
	//Avatar      string    `json:"avatar"`
}

func (u *User) SetPage(page Page) {
	u.Page = page
}

type Page struct {
	Slug      string `json:"slug" db:"page_slug" validate:"max=50"`
	IsPrivate bool   `json:"is_private" db:"page_is_private"`
}

type City struct {
	Id   IntId  `db:"city_id"`
	Name string `db:"name"`
}

type DisplayUserDto struct {
	Id      IntId  `json:"id" db:"user_id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Page
}

// DTO для регистрации пользователя
type SignupDto struct {
	Credentials `validate:"dive,required"`
	Name        string    `json:"name" validate:"required,max=50"`
	Surname     string    `json:"surname" validate:"required,max=50"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Gender      string    `json:"gender" validate:"required,oneof=м ж"`
	Interests   []string  `json:"interests"`
	CityId      IntId     `json:"cityId" validate:"required"`
}

func (dto SignupDto) ToUserWithPassword() *UserWithPassword {
	return &UserWithPassword{
		User: &User{
			Email:       dto.Email,
			Name:        dto.Name,
			Surname:     dto.Surname,
			DateOfBirth: dto.DateOfBirth,
			Gender:      dto.Gender,
			Interests:   dto.Interests,
			CityId:      dto.CityId,
		},
		Password: dto.Password,
	}
}

type UserWithPassword struct {
	*User
	Password string
}

// DTO для аутентификации
type Credentials struct {
	Email    string `json:"email" validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// идентификатор сессии
type Session string

func NewSession() Session {
	return Session(uuid.New().String())
}
