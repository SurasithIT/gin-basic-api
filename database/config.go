package database

type Config struct {
	Host         string `envconfig:"POSTGRES_HOST"`
	Username     string `envconfig:"POSTGRES_USERNAME"`
	Password     string `envconfig:"POSTGRES_PASSWORD"`
	DatabaseName string `envconfig:"POSTGRES_DATABASE"`
	Port         int    `envconfig:"POSTGRES_PORT" default:"5432"`
	Migrate      bool   `envconfig:"POSTGRES_MIGRATION" default:"false"`
}
