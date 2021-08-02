package models

type User struct {
	Username string `db:"username" json:"username" validate:"required,min=3,max=40,alphanum,nefield=Password"`
	Password string `db:"password_hash" json:"password" validate:"required,min=8"`
	Timezone string `db:"timezone" json:"timezone" validate:"required"`
}
