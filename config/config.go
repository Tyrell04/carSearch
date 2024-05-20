package config

type Config struct {
	Database
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
