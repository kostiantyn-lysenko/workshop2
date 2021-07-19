package models

type User struct {
	Username string `json:"username" validate:"required,min=3,max=40,alphanum,nefield=Password"`
	Password string `json:"password" validate:"required,min=8"`
	Timezone string `json:"timezone" validate:"required"`
}
