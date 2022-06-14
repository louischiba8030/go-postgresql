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
