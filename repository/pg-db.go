package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBDriver struct {
	config *Config
	logger *log.Logger
}

type Config struct {
	ConnectionString string `json:"connection_string"`
	Enabled          bool   `json:"enabled"`
	Port             string `json:"port"`
	Database         string `json:"database_name"`
}

func NewPostgresDB(config *Config, logger *log.Logger) *PostgresDBDriver {
	return &PostgresDBDriver{config, logger}
}

func (driver *PostgresDBDriver) ConnectPostgresDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf(driver.config.ConnectionString)

	//if os.Getenv("ENV") == "staging" || os.Getenv("ENV") == "production" {
	//	db, err = gorm.Open(postgres.Open(config.DBurl), &gorm.Config{})
	//} else {
	//	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//dsn := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	"localhost", driver.config.Port, "postgres", "1", driver.config.Database, "disable",
	//)
	//}
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return database, nil
}
