package usecase

import (
	"fmt"
	"math/rand"
)

func generateCode() (code string) {
	code = fmt.Sprintf("%04d", rand.Intn(10000))
	return
}
