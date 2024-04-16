package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Users []User

func AllUser(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Id: 1, Name: "majid", Username: "majid", Email: "majid@gmail.com"},
		User{Id: 2, Name: "samsul", Username: "samsul", Email: "samsul@gmail.com"},
	}

	fmt.Println("All API users")
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome API IDI Jakarta Pusat")
}

func handlerRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", AllUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlerRequests()
}
