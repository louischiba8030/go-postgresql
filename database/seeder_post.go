package database

import (
	"log"
	"context"

	"go-postgresql/ent"
	"go-postgresql/ent/migrate"
)

type Client struct {
	entclient *ent.Client
}

func PostSeeder (c *Client) {
	// Do Migration
	ctx := context.Background()
	if err := c.entclient.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

/*	// Add Post
	usr, err := Client.Debug().Post.
		Create().
		SetName("Miho Yamaguchi").
		SetAge(42).
		SetBloodtype("AB").
		SetOrigin("Sakado").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	log.Printf("user: %+v", usr)
	*/
}