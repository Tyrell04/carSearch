package database

import (
	"carSearch/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	*config.Config
	*sql.DB
}

func NewDatabase(cfg *config.Config) *Postgres {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(connStr)
		log.Fatal(err)
	}

	// Check if database is alive
	err = db.Ping()
	if err != nil {
		log.Println(connStr)
		log.Fatal(err)
	}

	return &Postgres{cfg, db}
}

func (postgres *Postgres) Close() {
	err := postgres.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
