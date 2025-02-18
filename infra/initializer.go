package infra

import (
	"github.com/joho/godotenv"

	"log"
)

func Initialize() {
	err := godotenv.Load() // .envファイルの読み込み
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
