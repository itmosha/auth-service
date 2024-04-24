package storage

import "errors"

var (
	ErrRegisterCodeNotExist    = errors.New("register code does not exist")
	ErrPhonenumberNotFound     = errors.New("user with provided phonenumber not found")
	ErrUidNotFound             = errors.New("user with provided uid not found")
	ErrPhonenumberAlreadyExist = errors.New("phonenumber already exists")
)
