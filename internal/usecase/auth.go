package usecase

type AuthStorageInterface interface{}

type AuthUsecase struct {
	store AuthStorageInterface
}

func NewAuthUsecase(authStorage AuthStorageInterface) *AuthUsecase {
	return &AuthUsecase{authStorage}
}
