package model

import "time"

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
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Interests   []string  `json:"interests"`
	City        string    `json:"city"`
}
