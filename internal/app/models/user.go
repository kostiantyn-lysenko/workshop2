package models

import "time"

type User struct {
	Username string
	Password string
	Token    Token
}

type Token struct {
	ExpiredAt time.Time
	Value     string
}
