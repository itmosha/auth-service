package usecase

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrRegistrationNotFinished = errors.New("this user did not finish registration")
	ErrAlreadyRegistered       = errors.New("this user is already registered")
	ErrWrongCodeProvided       = errors.New("wrong code provided")
)

func generateCode() (code string) {
	code = fmt.Sprintf("%04d", rand.Intn(10000))
	return
}
