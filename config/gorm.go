package config

import (
	"order-service/pkg/dotenv"

	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	const (
		IDLE_CONNECTION         = 10
		MAX_CONNECTION          = 100
		MAX_LIFETIME_CONNECTION = 300
	)

	username := dotenv.Getenv("MYSQL_USERNAME")
	password := dotenv.Getenv("MYSQL_PASSWORD")
	host := dotenv.Getenv("MYSQL_HOST")
	port := dotenv.Getenv("MYSQL_PORT")
	database := dotenv.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	connection.SetMaxIdleConns(IDLE_CONNECTION)
	connection.SetMaxOpenConns(MAX_CONNECTION)
	connection.SetConnMaxLifetime(time.Second * time.Duration(MAX_LIFETIME_CONNECTION))

	return db
}
