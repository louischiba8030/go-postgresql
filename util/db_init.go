package util

import (
	"context"
	"log"
	"fmt"

	"go-postgresql/config"
	"go-postgresql/ent"
	"go-postgresql/ent/migrate"

)

func dbInit () (context.Context, *ent.Client) {
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
