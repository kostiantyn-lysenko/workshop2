package models

type SignIn struct {
	Username string `json:"username" validate:"required,min=3,max=40,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=256,containsany=!@#?"`
}
