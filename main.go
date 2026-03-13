package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name       string
	Age        int
	Profession string
	IsAlive    bool
}

var currentUser User

func (u User) desc() string {
	s := fmt.Sprintf("Этот пользователь: %v, %v, %v", u.Name, u.Age, u.Profession)
	return s
}

func change(u *User) {
	u.Age = 30
	fmt.Println(u)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.Header)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func AuthorisationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Необходима авторизация!"))
}

func UserInformationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(currentUser.desc()))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("15:04:05 02.01.2006")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Тукущее время: " + currentTime))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Моё первое веб-приложение"))
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Версия 1.0.0"))
}

func main() {
	person := User{
		Name:       "Anna",
		Age:        20,
		Profession: "IT",
		IsAlive:    true,
	}

	change(&person)
	currentUser = person

	s := person.desc()
	fmt.Println(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/user", HelloHandler)
	mux.HandleFunc("/auth", AuthorisationHandler)
	mux.HandleFunc("/info", UserInformationHandler)
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/about", AboutHandler)
	mux.HandleFunc("/version", VersionHandler)

	log.Println("Сервер стартанул")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
