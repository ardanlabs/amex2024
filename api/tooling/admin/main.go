package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ardanlabs/service/business/sdk/migrate"
	"github.com/ardanlabs/service/business/sdk/sqldb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := sqldb.Config{
		User:         "postgres",
		Password:     "postgres",
		Host:         "database-service",
		Name:         "postgres",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		DisableTLS:   false,
	}

	if err := Migrate(cfg); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := Seed(cfg); err != nil {
		return fmt.Errorf("seed: %w", err)
	}

	return nil
}

func Migrate(cfg sqldb.Config) error {
	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}

func Seed(cfg sqldb.Config) error {
	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
