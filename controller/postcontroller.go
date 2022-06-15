package controller

import (
	"net/http"
	"context"
	"github.com/labstack/echo/v4"
	"go-postgresql/util"
	"go-postgresql/ent"
	"go-postgresql/ent/post"

	"log"
	"strconv"
//	"reflect"
)


func GetAllPosts(c echo.Context) (err error) {
	defer util.Ent().Close()
	ctx := context.Background()

	// Get all records and sort by ID
	posts, err := util.Ent().Debug().Post.
		Query().
		Order(ent.Asc(post.FieldID)).
		All(ctx)
	if err != nil {
		log.Fatalf("Failed to retreive posts: %v", err)
	}

	return c.JSON(http.StatusOK, posts) //*/ return nil
}

func AddPost (c echo.Context) error {
	defer util.Ent().Close()
	ctx := context.Background()

	p, _ := strconv.Atoi(c.FormValue("age"))
	age := int(p)

	// Generate ULID
	ulid := util.GenerateULID()

	newuser, err := util.Ent().Debug().Post.
		Create().
		SetID(ulid).
		SetName(c.FormValue("name")).
		SetAge(age).
		SetBloodtype(c.FormValue("bloodtype")).
		SetOrigin(c.FormValue("origin")).
		Save(ctx)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("user: %+v", newuser)

	return c.JSON(http.StatusOK, newuser)
}
