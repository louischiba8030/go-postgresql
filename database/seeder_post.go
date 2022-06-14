package database

import (
	"log"
	"fmt"
	"context"
	"os"
	"encoding/json"

	"go-postgresql/model"

	"go-postgresql/ent"
	"go-postgresql/util"
//	"go-postgresql/ent/migrate"
)

func InitialSeeder (ctx context.Context, client *ent.Client) error {
	// Get current directory
	var stcData []*model.Post
//	p, _ := os.Getwd()

	// Load local json file
	f, err := os.ReadFile("database/n46_members.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read stream from json data
	if err := json.Unmarshal([]byte(f), &stcData); err != nil {
		fmt.Println(err)
		return err
	}

	// Put each row and store into the SQL table
	for _, postRow := range stcData {
		// Generate ULID
		ulid := util.GenerateULID()
		fmt.Println(ulid)

		post := model.Post {
//			ID: ulid,
			Name: postRow.Name,
			Age: postRow.Age,
			Bloodtype: postRow.Bloodtype,
			Origin: postRow.Origin,
		}
		// Insert into PostgreSQL
		p, err := client.Debug().Post.
			Create().
			SetID(ulid).
			SetName(post.Name).
			SetAge(post.Age).
			SetBloodtype(post.Bloodtype).
			SetOrigin(post.Origin).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		log.Printf("user: %+v", p)
	} // End for-loop

	return err
}
