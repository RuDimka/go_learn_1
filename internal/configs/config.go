package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Port string
}

type DbConfig struct {
	DSN string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не смог загрузить ENV файл")
	}
	cfg := &Config{
		Db:   DbConfig{DSN: GetEnv("DSN", "postgres://taskuser:taskpass@localhost:5432/tasksdb?sslmode=disable")},
		Port: GetEnv("PORT", ":8080"),
	}

	return cfg
}

func GetEnv(k, d string) string {
	s := os.Getenv(k)
	if s == "" {
		log.Println("Ключ не нашёлся, применили default значение")
		return d
	}
	return s
}
