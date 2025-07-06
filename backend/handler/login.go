package handler

import (
	"net/http"

	"github.com/endr0id/locksmith/model"
	"github.com/endr0id/locksmith/pkg/auth"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/login", loginHandler)
}

func loginHandler(c echo.Context) error {
	req := new(model.Login)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Email != "admin@gmail.co.jp" || req.Password != "password" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := auth.GenerateJWT(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{ "message": token })
}