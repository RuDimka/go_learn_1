package main

import (
	"fmt"
	"log"
	"net/http"
	"test/app/internal/configs"
	"test/app/internal/database"
	"test/app/internal/handlers"
	"test/app/internal/service"
)

func main() {

	config := configs.NewConfig()
	fmt.Println(config)

	db, err := database.Connect(*config)
	if err != nil {
		log.Fatal("Ошибка подключения к БД", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatal("Ошибка миграции", err)
	}

	mux := handlers.NewRouter(service.NewUserService(db))

	log.Println("Сервер стартанул на порту: ", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, mux))
}
