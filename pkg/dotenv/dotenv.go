package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		return ""
	}

	return os.Getenv(key)
}
