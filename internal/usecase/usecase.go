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

type UsersMetaStorageInterface interface {
	Insert(ctx *context.Context, insertedUserMeta *entity.UserMeta) (createdUserMeta *entity.UserMeta, err error)
	SelectByPhonenumber(ctx *context.Context, phonenumber string) (userMeta *entity.UserMeta, err error)
	SelectByUid(ctx *context.Context, uid string) (userMeta *entity.UserMeta, err error)
	UpdateFields(ctx *context.Context, uid string, fields map[string]interface{}) (err error)
	DeleteUnregistered(ctx *context.Context, delta time.Duration) (err error)
}

type SessionsStorageInterface interface {
	Insert(ctx *context.Context, insertedSession *entity.Session) (createdSession *entity.Session, err error)
}

type CacheInterface interface {
	SetRegisterCode(ctx *context.Context, uid, code string) (err error)
	GetRegisterCode(ctx *context.Context, uid string) (code string, err error)
	DelRegisterCode(ctx *context.Context, uid string) (err error)
	SetLoginCode(ctx *context.Context, uid, code string) (err error)
	GetLoginCode(ctx *context.Context, uid string) (code string, err error)
	DelLoginCode(ctx *context.Context, uid string) (err error)
}

type Usecase struct {
	usersMetaStore UsersMetaStorageInterface
	sessionsStore  SessionsStorageInterface
	cache          CacheInterface
	jwtClient      *simplejwt.JWTClient
}

func NewUsecase(usersMetaStore UsersMetaStorageInterface, sessionsStore SessionsStorageInterface, cache CacheInterface, jwtClient *simplejwt.JWTClient) *Usecase {
	return &Usecase{usersMetaStore, sessionsStore, cache, jwtClient}
}

func (uc *Usecase) Register(ctx *context.Context, body *entity.RegisterBody) (userMeta *entity.UserMeta, err error) {
	userMeta, err = uc.usersMetaStore.SelectByPhonenumber(ctx, body.Phonenumber)
	if err == nil {
		if !userMeta.IsRegistered {
			err = ErrRegistrationNotFinished
		} else {
			err = ErrAlreadyRegistered
		}
		return
	} else if !errors.Is(err, storage.ErrUserMetaNotFound) {
		return
	}
	userMeta, err = uc.usersMetaStore.Insert(ctx, &entity.UserMeta{Phonenumber: body.Phonenumber})
	if err != nil {
		return
	}
	code := generateCode()
	fmt.Printf("Register code for uid %s: %s\n", userMeta.Uid, code)
	err = uc.cache.SetRegisterCode(ctx, userMeta.Uid, code)
	return
}

func (uc *Usecase) ConfirmRegister(ctx *context.Context, body *entity.ConfirmRegisterBody) (tokenPair *entity.TokenPair, err error) {
	userMeta, err := uc.usersMetaStore.SelectByUid(ctx, body.Uid)
	if err != nil {
		return
	} else if userMeta.IsRegistered {
		err = ErrAlreadyRegistered
		return
	}

	code, err := uc.cache.GetRegisterCode(ctx, userMeta.Uid)
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
	err = uc.usersMetaStore.UpdateFields(ctx, userMeta.Uid, map[string]interface{}{"is_registered": true})
	if err != nil {
		return
	}
	_ = uc.cache.DelRegisterCode(ctx, userMeta.Uid)
	atClaims := map[string]interface{}{"uid": userMeta.Uid, "phonenumber": userMeta.Phonenumber}
	rtClaims := map[string]interface{}{"uid": userMeta.Uid}
	tp, err := uc.jwtClient.CreateTokenPair(atClaims, rtClaims)
	if err != nil {
		return
	}
	tokenPair = &entity.TokenPair{
		AccessToken:  tp.AccessToken.Token,
		RefreshToken: tp.RefreshToken.Token,
	}
	_, err = uc.sessionsStore.Insert(ctx,
		&entity.Session{
			UserUid:   userMeta.Uid,
			Token:     tp.RefreshToken.Token,
			ExpiresAt: tp.RefreshToken.Exp,
			IssuedAt:  tp.RefreshToken.Iat,
		},
	)
	return
}

func (uc *Usecase) Login(ctx *context.Context, body *entity.LoginBody) (err error) {
	userMeta, err := uc.usersMetaStore.SelectByPhonenumber(ctx, body.Phonenumber)
	if err != nil {
		return
	} else if !userMeta.IsRegistered {
		err = ErrRegistrationNotFinished
		return
	}
	code := generateCode()
	fmt.Printf("Login code for uid %s: %s\n", userMeta.Uid, code)
	err = uc.cache.SetLoginCode(ctx, userMeta.Uid, code)
	return
}

func (uc *Usecase) ConfirmLogin(ctx *context.Context, body *entity.ConfirmLoginBody) (tokenPair *entity.TokenPair, err error) {
	userMeta, err := uc.usersMetaStore.SelectByUid(ctx, body.Uid)
	if err != nil {
		return
	} else if !userMeta.IsRegistered {
		err = ErrRegistrationNotFinished
		return
	}

	code, err := uc.cache.GetLoginCode(ctx, userMeta.Uid)
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
	_ = uc.cache.DelLoginCode(ctx, userMeta.Uid)
	atClaims := map[string]interface{}{"uid": userMeta.Uid, "phonenumber": userMeta.Phonenumber}
	rtClaims := map[string]interface{}{"uid": userMeta.Uid}
	tp, err := uc.jwtClient.CreateTokenPair(atClaims, rtClaims)
	if err != nil {
		return
	}
	tokenPair = &entity.TokenPair{
		AccessToken:  tp.AccessToken.Token,
		RefreshToken: tp.RefreshToken.Token,
	}
	_, err = uc.sessionsStore.Insert(ctx,
		&entity.Session{
			UserUid:   userMeta.Uid,
			Token:     tp.RefreshToken.Token,
			ExpiresAt: tp.RefreshToken.Exp,
			IssuedAt:  tp.RefreshToken.Iat,
		},
	)
	return
}
