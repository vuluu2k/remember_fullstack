package model

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
	Name     string    `db:"name" json:"name"`
	ImageUrl string    `db:"image_url" json:"image_url"`
	Website  string    `db:"website" json:"website"`
}
