package controller

import (
	"net/http"

	"github.com/itmosha/auth-service/pkg/logger"
)

type AuthUsecaseInterface interface{}

type AuthController struct {
	usecase AuthUsecaseInterface
	logger  *logger.Logger
}

func NewController(uc AuthUsecaseInterface, logger *logger.Logger) *AuthController {
	return &AuthController{uc, logger}
}

// @Title Register new user.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/register/ [post]
func (c *AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Confirm user registration.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/register/confirm/ [post]
func (c *AuthController) ConfirmRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Log in user.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/ [post]
func (c *AuthController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Confirm user login.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/confirm/ [post]
func (c *AuthController) ConfirmLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Refresh token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/refresh/ [post]
func (c *AuthController) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Revoke token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/revoke/ [post]
func (c *AuthController) Revoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}
