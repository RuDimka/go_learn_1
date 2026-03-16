package main

import (
	"fmt"
	"log"
	"net/http"
	"test/app/internal/configs"
	"test/app/internal/handlers"
	"test/app/internal/models"
)

func main() {
	person := models.User{
		Name:       "Anna",
		Age:        20,
		Profession: "IT",
		IsAlive:    true,
	}

	fmt.Println(person)

	config := configs.NewConfig()
	fmt.Println(config)

	mux := handlers.NewRouter(config.DBConn)

	log.Println("Сервер стартанул на порту: ", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, mux))
}
