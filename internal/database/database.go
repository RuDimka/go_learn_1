package database

import (
	"log"
	"test/app/internal/configs"
	"test/app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg configs.Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(cfg.Db.DSN), &gorm.Config{})
	if err != nil {
		log.Println("Не удалось подключиться к базе данных:", err)
		return nil, err
	}

	log.Println("Успешно подключились к базе данных")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Запущен процесс автомиграции")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Что-то не так")
		return err
	}
	log.Println("Миграция успешно завершена")
	return nil
}
