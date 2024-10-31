package main

import (
	"log"
	"net/http"

	"go-crud-backend/db"
	"go-crud-backend/router"
)

func main() {
    db.Init()
    r := router.InitRouter()
    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}