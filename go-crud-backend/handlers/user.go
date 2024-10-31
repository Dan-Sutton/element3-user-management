package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud-backend/db"
	"go-crud-backend/models"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    err := db.DB.QueryRow(
        "INSERT INTO users (first_name, surname, email, dob) VALUES ($1, $2, $3, $4) RETURNING id",
        user.FirstName, user.Surname, user.Email, user.DOB,
    ).Scan(&user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var user models.User
    err := db.DB.QueryRow("SELECT id, first_name, surname, email, dob FROM users WHERE id = $1", id).Scan(
        &user.ID, &user.FirstName, &user.Surname, &user.Email, &user.DOB,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var updatedUser models.User
    json.NewDecoder(r.Body).Decode(&updatedUser)
    updatedUser.ID = id

    _, err := db.DB.Exec(
        "UPDATE users SET first_name = $1, surname = $2, email = $3, dob = $4 WHERE id = $5",
        updatedUser.FirstName, updatedUser.Surname, updatedUser.Email, updatedUser.DOB, updatedUser.ID,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    _, err := db.DB.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}