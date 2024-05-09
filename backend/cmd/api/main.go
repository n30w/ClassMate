package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/n30w/Darkspace/internal/dal"
	"github.com/n30w/Darkspace/internal/domain"
	"log"
	"os"
)

const version = "1.0.0"

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var cfg config

	flag.IntVar(&cfg.port, "port", 6789, "API server port")
	flag.StringVar(
		&cfg.env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)

	// Database driver.
	flag.StringVar(&cfg.db.Dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	// Database configuration for connection settings.
	flag.IntVar(
		&cfg.db.MaxOpenConns, "db-max-open-conns", 25,
		"PostgreSQL max open connections",
	)
	flag.IntVar(
		&cfg.db.MaxIdleConns, "db-max-idle-conns", 25,
		"PostgreSQL max idle connections",
	)
	flag.StringVar(
		&cfg.db.MaxIdleTime, "db-max-idle-time", "15m",
		"PostgreSQL max connection idle time",
	)

	// Rate limiter configurations.
	flag.Float64Var(
		&cfg.limiter.rps,
		"limiter-rps",
		2,
		"Rate limiter maximum requests per second",
	)
	flag.IntVar(
		&cfg.limiter.burst,
		"limiter-burst",
		4,
		"Rate limiter maximum burst",
	)
	flag.BoolVar(
		&cfg.limiter.enabled,
		"limiter-enabled",
		true,
		"Enable rate limiter",
	)

	flag.Parse()

	logger := log.New(os.Stdout, "[DKSE] ", log.Ldate|log.Ltime)

	cfg.db.Driver = "postgres"

	// Set config database parameters via environment variables.
	// cfg.SetFromEnv()

	cfg.db.Dsn = os.Getenv("DB_DSN")
	cfg.db.Name = os.Getenv("DB_NAME")
	cfg.db.Username = os.Getenv("DB_USERNAME")
	cfg.db.Password = os.Getenv("DB_PASSWORD")
	cfg.db.Host = os.Getenv("DB_HOST")
	cfg.db.Port = os.Getenv("DB_PORT")
	cfg.db.SslMode = os.Getenv("DB_SSL_MODE")

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	store := dal.NewStore(db)
	excelStore := dal.NewExcelStore()

	app := &application{
		config:   cfg,
		logger:   logger,
		services: domain.NewServices(store, excelStore),
	}
	err = app.server()

	logger.Fatal(err)
}
