package config

type Config struct {
	DBPort             string `yaml:"db_port" env:"PORT-DB" env-default:"5432"`
	DBHost             string `yaml:"db_host" env:"HOST-DB" env-default:"localhost"`
	DBName             string `yaml:"db_name" env:"NAME-DB" env-default:"restapi_dev"`
	User               string `yaml:"user" env:"USER-DB" env-default:"user"`
	Password           string `yaml:"password" env:"PASSWORD-DB"`
	MigrationSource    string `yaml:"migration_source" env:"MIGRATION_SOURCE" env-default:"restapi_dev"`
	MigrationSchema    string `yaml:"migration_schema" env:"MIGRATION_SCHEMA" env-default:"file://pkg/db/migrations."`
	MigrationTableName string `yaml:"migration_table_name" env:"MIGRATION_TABLE_NAME" env-default:"msa_bank_client_cs"`
}
