package usecase

import (
	"errors"
)

var (
	ErrRegistrationNotFinished = errors.New("this user did not finish registration")
	ErrAlreadyRegistered       = errors.New("this user is already registered")
	ErrWrongCodeProvided       = errors.New("wrong code provided")
)
