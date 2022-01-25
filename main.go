package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Primeira letra deve ser em maiúsculo para exportação
type User struct {
	Id int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{{
		Id: 1,
		Name: "Gabriel",
	}})
}
