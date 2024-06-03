package main

import (
	"conn/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func basicAuth(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	fmt.Println(username, password, ok)

	if username != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unathorized")
		return
	}

	response := map[string]any{
		"data": "resource data",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	handler, err := models.NewHandler()
	if err != nil {
		log.Fatal("Error while connecting...", err)
	}

	log.Println("Connected successfully")
	http.HandleFunc("GET /resource", basicAuth)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	log.Println("Listening the :9000")
	http.ListenAndServe(":9000", nil)
}
