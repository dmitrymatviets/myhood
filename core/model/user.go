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
	City        City      `json:"city"`
	Page        Page      `json:"page"`
	Avatar      string    `json:"avatar"`
	FriendIds   []IntId   `json:"friends"`
}

func (u *User) AddFriend(friend *User) {
	for _, friendId := range u.FriendIds {
		if friendId == friend.Id {
			return
		}
	}
	u.FriendIds = append(u.FriendIds, friend.Id)
}

func (u *User) RemoveFriend(friend *User) {
	for id, friendId := range u.FriendIds {
		if friendId == friend.Id {
			u.FriendIds = append(u.FriendIds[:id], u.FriendIds[id+1:]...)
			return
		}
	}
}

func (u *User) SetPage(page Page) {
	u.Page = page
}

type Page struct {
	Slug      string
	IsPrivate bool
}

type City struct {
	Id   IntId
	Name string
}

type DisplayUserDto struct {
	Id      IntId  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Page    Page   `json:"page"`
	Avatar  string `json:"avatar"`
}

// DTO для регистрации пользователя
type SignupDto struct {
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Surname     string    `json:"surname" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Interests   []string  `json:"interests"`
	City        string    `json:"city" binding:"required"`
	AvatarFile  string    `json:"avatar"`
}

func (dto SignupDto) ToUser(city City) *User {
	return &User{
		Email:       dto.Email,
		Name:        dto.Name,
		Surname:     dto.Surname,
		DateOfBirth: dto.DateOfBirth,
		Gender:      dto.Gender,
		Interests:   dto.Interests,
		City:        city,
	}
}

// DTO для аутентификации
type Credentials struct {
	Email    string
	Password string
}

// идентификатор сессии
type Session string

func NewSession() Session {
	return Session(uuid.New().String())
}
