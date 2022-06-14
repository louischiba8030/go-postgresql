package database

import (
	"context"
	"log"
	"go-postgresql/ent"
)

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
