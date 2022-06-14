package main
import (
	"context"
	"fmt"
	"log"
//	"os"

	"go-postgresql/config"
	"go-postgresql/database"
	"go-postgresql/model"

//	"entgo.io/ent/examples/start/ent"
	_ "github.com/lib/pq"

	"go-postgresql/ent"
	"go-postgresql/ent/migrate"
)

func main () {
	//database.establish_connection()
	conf := config.Config
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbName, conf.DbPassword))
	if err != nil {
		log.Fatalln("Failed to get result", err)
	}
	defer client.Close()

	// Do migration
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	post := model.Post{
		Name: "Miho",
		Age: 42,
		Bloodtype: "B",
		Origin: "Sakado",
	}
	database.AddPost(ctx, client, &post)
	log.Print("ent sample done.")
}
