package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port         = 0
	DBConnection = ""
	SecretKey    []byte
)

func Load() {
	var err error

	if err = godotenv.Load(getRootPath() + ".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 3000
	}

	DBConnection = os.Getenv("DB_CONNECTION")
	if DBConnection == "" {
		DBConnection = "sqlite.db"
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Dir(b)
	rootPath = rootPath[:len(rootPath)-14]
	return rootPath
}
