package entity

import "time"

type AuthData struct {
	Uid          string    `json:"uid"`
	Phonenumber  string    `json:"phonenumber"`
	IsRegistered bool      `json:"is_registered"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
