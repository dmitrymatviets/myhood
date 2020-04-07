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
