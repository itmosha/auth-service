package storage

import "errors"

var (
	ErrPhonenumberNotFound     = errors.New("phonenumber not found")
	ErrPhonenumberAlreadyExist = errors.New("phonenumber already exists")
)
