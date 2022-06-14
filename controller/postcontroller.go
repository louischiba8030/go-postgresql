package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"go-postgresql/util"

//	"go-postgresql/model"
//	"strconv"
	"log"
)

func GetAllPosts(c echo.Context) error {
	ctx, client := util.DbInit()
	defer client.Close()

	posts, err := client.Debug().Post.
		Query().
		All(ctx)
	if err != nil {
		log.Fatalf("Failed to retreive posts: %v", err)
	}

	return c.JSON(http.StatusOK, posts)
}
