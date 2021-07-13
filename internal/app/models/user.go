package models

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    Token  `json:"token"`
}

type Token struct {
	ExpiredAt time.Time `json:"expired_at"`
	Value     string    `json:"value"`
}

func (u *User) Validate() error {
	return nil
}
