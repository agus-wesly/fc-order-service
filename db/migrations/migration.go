package migrations

import (
	"order-service/pkg/dotenv"

	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Start() {
	username := dotenv.Getenv("MYSQL_USERNAME")
	password := dotenv.Getenv("MYSQL_PASSWORD")
	host := dotenv.Getenv("MYSQL_HOST")
	port := dotenv.Getenv("MYSQL_PORT")
	database := dotenv.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	log.Println("Running migration...")
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to create mysql driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to init migrate: %v", err)
	}

	m.Steps(2)
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migration")
		} else {
			log.Fatal(err)
		}
	}

	log.Println("Done migration")
}
