package postgres

// conf is used as a default when running integration tests with docker-compose
var conf Config = Config{
	Host:     "localhost",
	Port:     5432,
	Database: "bookly",
	User:     "bookly",
	Password: "bookly",
}
