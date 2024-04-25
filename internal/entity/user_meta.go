package entity

import "time"

type UserMeta struct {
	Uid          string    `json:"uid" example:"9c884669-0dbf-497d-b94f-cfd196278d8f" db:"uid"`
	Phonenumber  string    `json:"phonenumber" example:"9009009090" db:"phonenumber"`
	IsRegistered bool      `json:"is_registered" example:"false" db:"is_registered"`
	CreatedAt    time.Time `json:"created_at" example:"2022-01-01T00:00:00Z" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z" db:"updated_at"`
}
