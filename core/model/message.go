package model

import "time"

type Message struct {
	Id           string
	Timestamp    time.Time
	AuthorUserId IntId
	Text         string
}

type Thread struct {
	ThreadId     IntId
	User         *DisplayUser
	Participants []*DisplayUser
	Messages     []*Message
}

type DisplayUser struct {
	Id      IntId  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Page    Page   `json:"page"`
	Avatar  string `json:"avatar"`
}
