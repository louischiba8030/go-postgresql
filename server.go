package main
import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-postgresql/controller"
)

func main () {
	router := initRouter()
	router.Logger.Fatal(router.Start(":8082"))
}

func initRouter() *echo.Echo {
	// Create instance
	e := echo.New()

	// Set middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang World!")
	})
	e.GET("/api/posts/", controller.GetAllPosts)
	e.POST("/api/posts/create", controller.AddPost)

	return e
}

/*
//database.establish_connection()
	conf := config.Config
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbName, conf.DbPassword))
	if err != nil {
		log.Fatalln("Failed to get result", err)
	}
	defer client.Close()

	// Drop existing table 'posts'

	// Do migration
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	// Call seeder (register dummy 10 posts)
	//database.InitialSeeder(ctx, client)
*/
