package model

// DTO для аутентификации
type Credentials struct {
	Email    string
	Password string
}

// идентификатор сессии
type Session string
