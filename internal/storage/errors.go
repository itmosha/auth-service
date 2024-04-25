package storage

import "errors"

var (
	// Redis errors
	ErrRegisterCodeNotExist = errors.New("register code does not exist")

	// Postgres errors
	ErrUidNotFound          = errors.New("user meta with provided uid was not found")
	ErrUserMetaAlreadyExist = errors.New("phonenumber already exists")
	ErrUserMetaNotFound     = errors.New("user with provided phonenumber was not found")
)
