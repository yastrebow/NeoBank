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
)

func Init(dbAddress, schemaName string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   schemaName,
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migration(
	dbAddress,
	migrationSource,
	migrationTableName string) {

	fullDBUrl := fmt.Sprintf("%s&x-migrations-table=%s",
		dbAddress, migrationTableName)

	m, err := migrate.New(migrationSource, fullDBUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
