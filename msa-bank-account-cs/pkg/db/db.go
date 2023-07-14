package db

import (
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type ConfigDatabase struct {
	Port     string `yaml:"port" env:"PORT-DB" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST-DB" env-default:"localhost"`
	User     string `yaml:"user" env:"USER-DB" env-default:"user"`
	Password string `yaml:"password" env:"PASSWORD-DB"`
}

func Init() *gorm.DB {
	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln("Cannot read config", err)
	}
	dbURL := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/restapi_dev?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "msa_bank_account_cs_schema.",
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func Migration() {
	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln("Cannot read config", err)
	}
	m, err := migrate.New(
		"file://pkg/db/migrations",
		"postgres://"+cfg.User+":"+cfg.Password+"@"+cfg.Host+":"+cfg.Port+"/restapi_dev?sslmode=disable&x-migrations-table=msa-bank-account-cs")
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
