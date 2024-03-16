package main

import (
	"fmt"
	"geo-data/config"
	"geo-data/routes"
	"net/http"
)

func init() {
	_, err := config.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	}
}

func main() {
	http.HandleFunc("/register", routes.Register)
	http.HandleFunc("/login", routes.Login)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
