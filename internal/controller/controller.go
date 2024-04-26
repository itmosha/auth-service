package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/itmosha/auth-service/internal/entity"
	"github.com/itmosha/auth-service/internal/storage"
	"github.com/itmosha/auth-service/internal/usecase"
	"github.com/itmosha/auth-service/pkg/logger"
)

type UsecaseInterface interface {
	Register(ctx *context.Context, body *entity.RegisterBody) (userMeta *entity.UserMeta, err error)
	ConfirmRegister(ctx *context.Context, body *entity.ConfirmRegisterBody) (tokenPair *entity.TokenPair, err error)
	Login(ctx *context.Context, body *entity.LoginBody) (err error)
}

type Controller struct {
	uc        UsecaseInterface
	validator *validator.Validate
	logger    *logger.Logger
}

func NewController(uc UsecaseInterface, logger *logger.Logger) *Controller {
	return &Controller{uc: uc, validator: validator.New(), logger: logger}
}

// @Title Register new user.
// @Descriptiomn Register a new user using a phonenumber.
// @Param body body entity.RegisterBody true "Registration body"
// @Success 201 object entity.UserMeta "Successful registration, user meta in response body"
// @Failure 400 object errorResponseBody "Invalid request body"
// @Failure 409 object errorResponseBody "User already registered"
// @Failure 422 object errorResponseBody "User registration not finished"
// @Failure 500 object errorResponseBody "Internal server error"
// @Resource Auth
// @Route /api/register/ [post]
func (c *Controller) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		body, err := readBodyToStruct(r, &entity.RegisterBody{})
		if err != nil {
			c.responseWithError(w, r, http.StatusBadRequest, err)
			return
		}
		err = c.validator.StructCtx(ctx, body)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			c.responseWithError(w, r, http.StatusBadRequest, errors)
			return
		}
		userMeta, err := c.uc.Register(&ctx, body)
		if err != nil {
			if errors.Is(err, usecase.ErrAlreadyRegistered) {
				c.responseWithError(w, r, http.StatusConflict, err)
			} else if errors.Is(err, usecase.ErrRegistrationNotFinished) {
				c.responseWithError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				c.responseWithError(w, r, http.StatusInternalServerError, err)
			}
			return
		}
		c.responseWithSuccess(w, r, http.StatusCreated, userMeta)
	}
}

// @Title Confirm user registration.
// @Descriptiomn Confirm user registration by uid.
// @Param body body entity.ConfirmRegisterBody true "Confirm registration body"
// @Success 200 object entity.TokenPair "Successful confirmation, token pair in response body"
// @Failure 400 object errorResponseBody "Invalid request body or user does not exist"
// @Failure 409 object errorResponseBody "User already registered"
// @Failure 422 object errorResponseBody "Wrong confirmation code provided"
// @Failure 500 object errorResponseBody "Internal server error"
// @Resource Auth
// @Route /api/register/confirm/ [post]
func (c *Controller) ConfirmRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		body, err := readBodyToStruct(r, &entity.ConfirmRegisterBody{})
		if err != nil {
			c.responseWithError(w, r, http.StatusBadRequest, err)
			return
		}
		err = c.validator.StructCtx(ctx, body)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			c.responseWithError(w, r, http.StatusBadRequest, errors)
			return
		}
		tokenPair, err := c.uc.ConfirmRegister(&ctx, body)
		if err != nil {
			if errors.Is(err, usecase.ErrAlreadyRegistered) {
				c.responseWithError(w, r, http.StatusConflict, err)
			} else if errors.Is(err, storage.ErrUserMetaNotFound) {
				c.responseWithError(w, r, http.StatusBadRequest, err)
			} else if errors.Is(err, usecase.ErrWrongCodeProvided) {
				c.responseWithError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				c.responseWithError(w, r, http.StatusInternalServerError, err)
			}
			return
		}
		c.responseWithSuccess(w, r, http.StatusOK, tokenPair)
	}
}

// @Title Log in.
// @Descriptiomn Log in a user by phonenumber.
// @Param body body entity.LoginBody true "Login body"
// @Success 204 "Successful confirmation, no content in response body"
// @Failure 400 object errorResponseBody "Invalid request body or user does not exist"
// @Failure 422 object errorResponseBody "User registration not finished"
// @Failure 500 object errorResponseBody "Internal server error"
// @Resource Auth
// @Route /api/login/ [post]
func (c *Controller) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		body, err := readBodyToStruct(r, &entity.LoginBody{})
		if err != nil {
			c.responseWithError(w, r, http.StatusBadRequest, err)
			return
		}
		err = c.validator.StructCtx(ctx, body)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			c.responseWithError(w, r, http.StatusBadRequest, errors)
			return
		}
		err = c.uc.Login(&ctx, body)
		if err != nil {
			if errors.Is(err, usecase.ErrRegistrationNotFinished) {
				c.responseWithError(w, r, http.StatusUnprocessableEntity, err)
			} else if errors.Is(err, storage.ErrUserMetaNotFound) {
				c.responseWithError(w, r, http.StatusBadRequest, err)
			} else {
				c.responseWithError(w, r, http.StatusInternalServerError, err)
			}
			return
		}
		c.responseWithSuccess(w, r, http.StatusNoContent, nil)
	}
}

// @Title Confirm user login.
// @Failure 501 {object} errorResponseBody
// @Resource Auth
// @Route /api/login/confirm/ [post]
func (c *Controller) ConfirmLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.responseWithError(w, r, http.StatusNotImplemented, ErrNotImplemented)
	}
}

// @Title Refresh token pair.
// @Failure 501 {object} errorResponseBody
// @Resource Auth
// @Route /api/refresh/ [post]
func (c *Controller) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.responseWithError(w, r, http.StatusNotImplemented, ErrNotImplemented)
	}
}

// @Title Revoke token pair.
// @Failure 501 {object} errorResponseBody
// @Resource Auth
// @Route /api/revoke/ [post]
func (c *Controller) Revoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.responseWithError(w, r, http.StatusNotImplemented, ErrNotImplemented)
	}
}

func (c *Controller) responseWithSuccess(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
	c.logger.LogRequest(r, statusCode, nil)
}

func (c *Controller) responseWithError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if statusCode >= 500 {
		json.NewEncoder(w).Encode(errorResponseBody{Message: ErrServerError.Error()})
	} else {
		json.NewEncoder(w).Encode(errorResponseBody{Message: err.Error()})
	}
	c.logger.LogRequest(r, statusCode, err)
}
