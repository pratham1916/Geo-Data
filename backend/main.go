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

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	http.HandleFunc("/register", corsMiddleware(routes.Register))
	http.HandleFunc("/login", corsMiddleware(routes.Login))
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
