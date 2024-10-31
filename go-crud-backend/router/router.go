package router

import (
	"go-crud-backend/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    return r
}