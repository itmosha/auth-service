package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/itmosha/auth-service/internal/entity"
	storage "github.com/itmosha/auth-service/internal/storage/postgres"
)

type StorageInterface interface {
	Insert(ctx *context.Context, insertedAuthData *entity.AuthData) (createdAuthData *entity.AuthData, err error)
	SelectByPhonenumber(ctx *context.Context, phonenumber string) (authData *entity.AuthData, err error)
	DeleteUnregistered(ctx *context.Context, delta time.Duration) (err error)
}

type CacheInterface interface {
	SetRegisterCode(ctx *context.Context, uid, code string) (err error)
	GetRegisterCode(ctx *context.Context, uid string) (code string, err error)
	DelRegisterCode(ctx *context.Context, uid string) (err error)
}

type Usecase struct {
	store StorageInterface
	cache CacheInterface
}

func NewUsecase(storage StorageInterface, cache CacheInterface) *Usecase {
	return &Usecase{storage, cache}
}

func (uc *Usecase) Register(ctx *context.Context, body *entity.RegisterBody) (authData *entity.AuthData, err error) {
	authData, err = uc.store.SelectByPhonenumber(ctx, body.Phonenumber)
	if err == nil {
		if !authData.IsRegistered {
			err = ErrRegistrationNotFinished
		} else {
			err = ErrAlreadyRegistered
		}
		return
	} else if !errors.Is(err, storage.ErrPhonenumberNotFound) {
		return
	}
	authData, err = uc.store.Insert(ctx, &entity.AuthData{Phonenumber: body.Phonenumber})
	if err != nil {
		return
	}
	code := generateCode()
	fmt.Printf("Code for uid %s: %s\n", authData.Uid, code)
	err = uc.cache.SetRegisterCode(ctx, authData.Uid, code)
	return
}
