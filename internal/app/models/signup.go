package models

type SignUp struct {
	Username       string `json:"username" validate:"required,min=3,max=40,alphanum,nefield=Password"`
	Password       string `json:"password" validate:"required,min=8,max=256,containsany=!@#?"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=Password"`
}
