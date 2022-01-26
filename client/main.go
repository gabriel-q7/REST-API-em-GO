package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// := atribuição dinamica do tipo
	resp, err := http.Get("http://localhost:8080/users")
	// problema no endpoint
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 200 = sucesso
	// endpoint funcionou mas não retornou o que precisava
	if resp.StatusCode != 200 {
		fmt.Println("Not sucess", resp.StatusCode)
		return
	}

	//Leitura da requisição
	body, err := io.ReadAll(resp.Body)

	var response []User
	json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Error to recover data", err.Error())
		return
	}

	fmt.Println(response)
}
