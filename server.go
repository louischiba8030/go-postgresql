package main
import (
	"context"
	"fmt"
	"log"
//	"os"

	"go-postgresql/config"
//	"go-postgresql/database"

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

	AddPost(ctx, client, "foo", 23, "A", "Chiba")
	log.Print("ent sample done.")
}

func AddPost (ctx context.Context, client *ent.Client, name string, age int, bloodtype string, origin string) (*ent.Post, error) {
	post, err := client.Debug().Post.
		Create().
		SetName(name).
		SetAge(age).
		SetBloodtype(bloodtype).
		SetOrigin(origin).
		Save(ctx)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("user: %+v", post)

	return post, nil
}
