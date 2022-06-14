package main

import (
	"net/http"
//	"fmt"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//
//	"go-postgresql/model"
	"go-postgresql/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func main () {
	// Create instance
	e := echo.New()

	// Set middleware
	e.Use(middleware.CORS())
//	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang World!")
	})
	e.GET("/api/posts/", controller.GetAllPosts)
	e.POST("/api/posts/create", controller.CreatePost)
/*
	e.GET("/api/posts/:id", controller.GetPost)
	e.PUT("/api/posts/:id", controller.UpdatePost)
	e.DELETE("/api/posts/:id", controller.DeletePost)
*/
	// Launch server at port 8003
	e.Logger.Fatal(e.Start(":8004"))

}
