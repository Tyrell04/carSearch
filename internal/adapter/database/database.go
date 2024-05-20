package database

import "carSearch/config"

type Database interface {
	NewDatabase(config *config.Config) error
	Close()
}
