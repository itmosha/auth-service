package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/itmosha/auth-service/internal/entity"
	"github.com/itmosha/auth-service/internal/usecase"
	"github.com/itmosha/auth-service/pkg/logger"
)

type UsecaseInterface interface {
	Register(ctx *context.Context, body *entity.RegisterBody) (authData *entity.AuthData, err error)
	ConfirmRegister(ctx *context.Context, body *entity.ConfirmRegisterBody) (tokenPair *entity.TokenPair, err error)
}

type Controller struct {
	uc        UsecaseInterface
	validator *validator.Validate
}

func NewController(uc UsecaseInterface, logger *logger.Logger) *Controller {
	return &Controller{uc, validator.New()}
}

// @Title Register new user.
// @Descriptiomn Register a new user using a phonenumber.
// @Param body body entity.RegisterBody true "Registration body"
// @Success 201 object entity.AuthData "Successful registration"
// @Failure 400 object ErrorResponseBody "Invalid request body"
// @Failure 409 object ErrorResponseBody "User already registered"
// @Failure 422 object ErrorResponseBody "User registration not finished"
// @Failure 500 object ErrorResponseBody "Internal server error"
// @Resource Auth
// @Route /api/register/ [post]
func (c *Controller) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		body, err := readBodyToStruct(r, &entity.RegisterBody{})
		if err != nil {
			ResponseWithError(w, r, http.StatusBadRequest, err)
			return
		}
		err = c.validator.StructCtx(ctx, body)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			ResponseWithError(w, r, http.StatusBadRequest, errors)
			return
		}
		authData, err := c.uc.Register(&ctx, body)
		if err != nil {
			if errors.Is(err, usecase.ErrAlreadyRegistered) {
				ResponseWithError(w, r, http.StatusConflict, err)
			} else if errors.Is(err, usecase.ErrRegistrationNotFinished) {
				ResponseWithError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				ResponseWithError(w, r, http.StatusInternalServerError, err)
			}
			return
		}
		ResponseWithSuccess(w, r, http.StatusCreated, authData)
	}
}

// @Title Confirm user registration.
// @Descriptiomn Confirm user registration by uid.
// @Param body body entity.ConfirmRegisterBody true "Confirm registration body"
// @Success 200 object entity.TokenPair "Successful confirmation"
// @Failure 400 object ErrorResponseBody "Invalid request body"
// @Failure 409 object ErrorResponseBody "User already registered"
// @Failure 422 object ErrorResponseBody "Wrong confirmation code provided"
// @Failure 500 object ErrorResponseBody "Internal server error"
// @Resource Auth
// @Route /api/register/confirm/ [post]
func (c *Controller) ConfirmRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		body, err := readBodyToStruct(r, &entity.ConfirmRegisterBody{})
		if err != nil {
			ResponseWithError(w, r, http.StatusBadRequest, err)
			return
		}
		err = c.validator.StructCtx(ctx, body)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			ResponseWithError(w, r, http.StatusBadRequest, errors)
			return
		}
		tokenPair, err := c.uc.ConfirmRegister(&ctx, body)
		if err != nil {
			if errors.Is(err, usecase.ErrAlreadyRegistered) {
				ResponseWithError(w, r, http.StatusConflict, err)
			} else if errors.Is(err, usecase.ErrWrongCodeProvided) {
				ResponseWithError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				ResponseWithError(w, r, http.StatusInternalServerError, err)
			}
			return
		}
		ResponseWithSuccess(w, r, http.StatusOK, tokenPair)
	}
}

// @Title Log in user.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/ [post]
func (c *Controller) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, r, http.StatusNotImplemented, nil)
	}
}

// @Title Confirm user login.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/confirm/ [post]
func (c *Controller) ConfirmLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, r, http.StatusNotImplemented, nil)
	}
}

// @Title Refresh token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/refresh/ [post]
func (c *Controller) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, r, http.StatusNotImplemented, nil)
	}
}

// @Title Revoke token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/revoke/ [post]
func (c *Controller) Revoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, r, http.StatusNotImplemented, nil)
	}
}
