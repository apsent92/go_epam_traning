package main

import (
	"fmt"
	"net/http"
	"restApi/internal/apiserver"
)

func main() {
	server := apiserver.APIServer{}
	server.Start()
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")

	server := apiserver.APIServer{}
	server.Start()
}


