package models

import (
	"conn/postgres"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler() (*Handler, error) {
	db, err := postgres.Initialize()

	if err != nil {
		log.Fatal("Error while connecting...", err)
	}

	return &Handler{DB: db}, nil
}

type Reg struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var check bool 
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method"))
		return 
	}
	var login Reg
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Fatal("error while decoding...", err)
	}

	rows, err := h.DB.Query("select * from register")
	if err != nil {
		log.Fatal(
			"Error while getting info from database...", err)
	}

	for rows.Next() {
		var log Reg
		if err = rows.Scan(&log.ID, &log.Name, &log.Email, &log.Password); err != nil {
			fmt.Printf("Error")
			return 
		}
		if login.Email == log.Email && login.Password == log.Password{
			check = true 
			break
		}
	}
	if check {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Successfully logged in"))
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
	

}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("wrong method"))
		return
	}
	var user Reg

	json.NewDecoder(r.Body).Decode(&user)

	_, err := h.DB.Exec("insert into register(name, email, password) values($1, $2, $3)", user.Name, user.Email, user.Password)

	if err != nil {
		log.Fatal("Error while inserting data to the table...", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successfully registreted"))

}
