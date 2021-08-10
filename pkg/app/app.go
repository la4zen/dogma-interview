package app

import (
	"github.com/la4zen/dogma-intwrview/pkg/service"
	"github.com/la4zen/dogma-intwrview/pkg/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	store := store.NewStore()
	service := service.NewService(store)
	// на самом деле я не могу сравнивать echo с gin,
	// хотя echo как фреймворк показался мне попроще.
	e := echo.New()
	e.Use(middleware.Logger())
	// e.Use(middleware.CORS()) CORS, CORSWithConfig
	e.POST("/users/create", service.CreateUser)
	e.PUT("/users/edit", service.EditUser)
	e.GET("/users", service.GetUsers)
	e.DELETE("/users/delete", service.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
