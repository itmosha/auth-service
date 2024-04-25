package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/itmosha/auth-service/internal/entity"
	"github.com/itmosha/auth-service/internal/storage"
	"github.com/itmosha/simplejwt"
)

type StorageInterface interface {
	Insert(ctx *context.Context, insertedAuthData *entity.AuthData) (createdAuthData *entity.AuthData, err error)
	SelectByPhonenumber(ctx *context.Context, phonenumber string) (authData *entity.AuthData, err error)
	SelectByUid(ctx *context.Context, uid string) (authData *entity.AuthData, err error)
	UpdateFields(ctx *context.Context, uid string, fields map[string]interface{}) (err error)
	DeleteUnregistered(ctx *context.Context, delta time.Duration) (err error)
}

type CacheInterface interface {
	SetRegisterCode(ctx *context.Context, uid, code string) (err error)
	GetRegisterCode(ctx *context.Context, uid string) (code string, err error)
	DelRegisterCode(ctx *context.Context, uid string) (err error)
}

type Usecase struct {
	store     StorageInterface
	cache     CacheInterface
	jwtClient *simplejwt.JWTClient
}

func NewUsecase(storage StorageInterface, cache CacheInterface, jwtClient *simplejwt.JWTClient) *Usecase {
	return &Usecase{storage, cache, jwtClient}
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

func (uc *Usecase) ConfirmRegister(ctx *context.Context, body *entity.ConfirmRegisterBody) (tokenPair *entity.TokenPair, err error) {
	authData, err := uc.store.SelectByUid(ctx, body.Uid)
	if err != nil {
		return
	} else if authData.IsRegistered {
		err = ErrAlreadyRegistered
		return
	}

	code, err := uc.cache.GetRegisterCode(ctx, authData.Uid)
	if errors.Is(err, storage.ErrRegisterCodeNotExist) {
		err = ErrWrongCodeProvided
		return
	} else if err != nil {
		return
	}
	if body.Code != code {
		err = ErrWrongCodeProvided
		return
	}
	err = uc.store.UpdateFields(ctx, authData.Uid, map[string]interface{}{"is_registered": true})
	if err != nil {
		return
	}
	_ = uc.cache.DelRegisterCode(ctx, authData.Uid)
	atClaims := map[string]interface{}{"uid": authData.Uid, "phonenumber": authData.Phonenumber}
	rtClaims := map[string]interface{}{"uid": authData.Uid}
	tp, err := uc.jwtClient.CreateTokenPair(atClaims, rtClaims)
	if err != nil {
		return
	}
	tokenPair = &entity.TokenPair{
		AccessToken:  tp.AccessToken,
		RefreshToken: tp.RefreshToken,
	}
	return
}
