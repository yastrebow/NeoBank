package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"msa-bank-client/internal/config"
)

func Init(cfg config.Config) *gorm.DB {
	dbAddress := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.MigrationSchema,
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migration(cfg config.Config) {

	dbAddress := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	m, err := migrate.New(cfg.MigrationSource, dbAddress)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
