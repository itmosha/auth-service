package usecase

type StorageInterface interface{}

type Usecase struct {
	store StorageInterface
}

func NewUsecase(storage StorageInterface) *Usecase {
	return &Usecase{storage}
}
