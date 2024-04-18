package controller

import (
	"net/http"

	"github.com/itmosha/auth-service/pkg/logger"
)

type UsecaseInterface interface{}

type Controller struct {
	usecase UsecaseInterface
	logger  *logger.Logger
}

func NewController(uc UsecaseInterface, logger *logger.Logger) *Controller {
	return &Controller{uc, logger}
}

// @Title Register new user.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/register/ [post]
func (c *Controller) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Confirm user registration.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/register/confirm/ [post]
func (c *Controller) ConfirmRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Log in user.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/ [post]
func (c *Controller) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Confirm user login.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/login/confirm/ [post]
func (c *Controller) ConfirmLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Refresh token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/refresh/ [post]
func (c *Controller) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}

// @Title Revoke token pair.
// @Failure 501 {object} ErrorResponseBody
// @Resource Auth
// @Route /api/revoke/ [post]
func (c *Controller) Revoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ResponseWithError(w, http.StatusNotImplemented, ErrServerError)
		c.logger.Log(r, http.StatusNotImplemented, nil)
	}
}
