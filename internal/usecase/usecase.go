package usecase

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrRegistrationNotFinished = errors.New("this user did not finish registration")
	ErrAlreadyRegistered       = errors.New("this user is already registered")
)

func generateCode() (code string) {
	code = fmt.Sprintf("%06d", rand.Intn(1000000))
	return
}
