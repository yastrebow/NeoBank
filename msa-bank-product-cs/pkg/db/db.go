package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "msa_bank_product_cs_schema.", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                          // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migration() {
	m, err := migrate.New(
		"file://pkg/db/migrations",
		"postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable&x-migrations-table=msa_bank_product_cs")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
