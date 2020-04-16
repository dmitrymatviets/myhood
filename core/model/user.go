package model

import (
	"github.com/google/uuid"
	"time"
)

type IntId int64

type User struct {
	Id          IntId     `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Interests   []string  `json:"interests"`
	CityId      IntId     `json:"cityId"`
	Page        Page      `json:"page"`
	//Avatar      string    `json:"avatar"`
}

func (u *User) SetPage(page Page) {
	u.Page = page
}

type Page struct {
	Slug      string
	IsPrivate bool
}

type City struct {
	Id   IntId  `db:"city_id"`
	Name string `db:"name"`
}

type DisplayUserDto struct {
	Id      IntId  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Page    Page   `json:"page"`
}

// DTO для регистрации пользователя
type SignupDto struct {
	Credentials `validate:"dive,required"`
	Name        string    `json:"name" validate:"required"`
	Surname     string    `json:"surname" validate:"required"`
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
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// идентификатор сессии
type Session string

func NewSession() Session {
	return Session(uuid.New().String())
}
