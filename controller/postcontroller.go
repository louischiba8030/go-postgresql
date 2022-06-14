package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"go-postgresql/model"
	"strconv"
	"fmt"
)

func GetAllPosts(c echo.Context) error {
	m := []model.Post{}
	model.DB.Find(&m)

	return c.JSON(http.StatusOK, m)
}

func CreatePost(c echo.Context) (err error) {
	p, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(p)

	m := model.Post {
		Name: c.FormValue("name"),
		Age: age,
		Bloodtype: c.FormValue("bloodtype"),
		Origin: c.FormValue("origin"),
	}

	fmt.Print("add post: ", m)
	m.Create()
	
	return c.JSON(http.StatusOK, m)
}
