package util

import (
//	"context"
	"log"
	"fmt"
	"database/sql"
	"go-postgresql/config"
	"go-postgresql/ent"
//	_ "github.com/lib/pq"
	"entgo.io/ent/dialect"
entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib" // https://entgo.io/ja/docs/sql-integration/
//	"go-postgresql/ent/migrate"

)

var entClient *ent.Client

func Ent () *ent.Client {
	return entClient
}

func DbInit_pgx () error {
	conf := config.Config
	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbName, conf.DbPassword))
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
//	return ent.NewClient(ent.Driver(drv))
	entClient = ent.NewClient(ent.Driver(drv))
	return nil
}

func DbInit () error { // (context.Context, *ent.Client) {
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
