package constant

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(paths ...string) {
	if len(paths) == 0 {
		paths = append(paths, ".env")
	}
	if err := godotenv.Load(paths...); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}
