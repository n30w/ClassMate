package dal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Set up a connection to a database.

func NewDBConfig() DBConfig {
	return DBConfig{
		Driver:       "postgres",
		MaxOpenConns: 25,
		MaxIdleConns: 25,
		MaxIdleTime:  "15m",
	}
}

type DBConfig struct {
	// Database driver and DataSourceName
	Driver string
	Dsn    string

	// Database parameters, similar to .env file variables.
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	SslMode  string

	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func (d DBConfig) SetFromEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d.Dsn = os.Getenv("DB_DSN")
	d.Name = os.Getenv("DB_NAME")
	d.Username = os.Getenv("DB_USERNAME")
	d.Password = os.Getenv("DB_PASSWORD")
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.SslMode = os.Getenv("DB_SSL_MODE")
}

// createDataSourceName creates the dataSourceName parameter of the
// sql.Open function.
func (d DBConfig) CreateDataSourceName() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		d.Port,
		d.Username,
		d.Password,
		d.Name,
		d.SslMode,
	)
}
