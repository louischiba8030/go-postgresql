package util

import (
//	"context"
	"log"
	"fmt"

	"go-postgresql/config"
	"go-postgresql/ent"
//	"go-postgresql/ent/migrate"

)

var entClient *ent.Client

func Ent () *ent.Client {
	return entClient
}

func DbInit () *ent.Client { // (context.Context, *ent.Client) {
	conf := config.Config
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbName, conf.DbPassword))
	if err != nil {
		log.Fatalln("Failed to get result", err)
	}

	// Drop existing table 'posts'

	// Do migration
//	ctx := context.Background()
	//if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
//		log.Fatalf("failed printing schema changes: %v", err)
//	}

	// Call seeder (register dummy 10 posts)

	entClient = client
	return nil
}
