package models

import (
	"time"
)

type User struct {
	Username string `json:"username" validate:"required,min=3,max=40,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=256,containsany=!@#?"`
	Token    Token  `json:"token" validate:"required"`
}

type Token struct {
	ExpiredAt time.Time `json:"expired_at" validate:"required"`
	Value     string    `json:"value" validate:"required"`
}
