package service

import (
	"github.com/la4zen/dogma-intwrview/pkg/models"
	"github.com/labstack/echo"
)

func (s *Service) GetUsers(c echo.Context) error {
	users, err := s.Store.GetUsers()
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"users": users,
	})
}

func (s *Service) EditUser(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.ValidateEdit(); err != nil {
		return c.String(400, "id required")
	}
	if err := s.Store.EditUser(user); err != nil {
		return c.String(500, err.Error())
	}
	return c.NoContent(200)
}

func (s *Service) CreateUser(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.ValidateCreate(); err != nil {
		return c.String(400, err.Error())
	}
	if err := s.Store.CreateUser(user); err != nil {
		return c.String(500, err.Error())
	}
	return c.NoContent(200)
}

func (s *Service) DeleteUser(c echo.Context) error {
	user := &models.User{}
	c.Bind(user)
	if err := user.Validate(); err != nil {
		return c.String(400, "login or id required")
	}
	if err := s.Store.DeleteUser(user); err != nil {
		return c.String(500, err.Error())
	}
	return c.NoContent(200)
}
