package database

import (
	"context"
	"log"
	"fmt"

	"go-postgresql/config"
	"go-postgresql/ent"
	"go-postgresql/ent/migrate"
	"go-postgresql/model"
)

func db_init () (context.Context, *ent.Client) {
	conf := config.Config
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbName, conf.DbPassword))
	if err != nil {
		log.Fatalln("Failed to get result", err)
	}
	defer client.Close()

	// Drop existing table 'posts'

	// Do migration
	ctx_rr := context.Background()
	if err := client.Schema.Create(ctx_rr, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	// Call seeder (register dummy 10 posts)
	//database.InitialSeeder(ctx, client)
	return ctx_rr, client
}

func AddPost (ctx context.Context, client *ent.Client, dPointer *model.Post) (*ent.Post, error) {
	d := *dPointer

	p, err := client.Debug().Post.
		Create().
		SetName(d.Name).
		SetAge(d.Age).
		SetBloodtype(d.Bloodtype).
		SetOrigin(d.Origin).
		Save(ctx)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("user: %+v", p)

	return p, nil
}
