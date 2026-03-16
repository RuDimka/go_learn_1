package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db     DbConfig
	Port   string
	DBConn *gorm.DB
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

	db, err := gorm.Open(postgres.Open(cfg.Db.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	cfg.DBConn = db
	log.Println("Успешно подключились к базе данных")

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
