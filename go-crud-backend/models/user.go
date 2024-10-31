package models

import "time"

type User struct {
    ID        int       `json:"id"`
    FirstName string    `json:"first_name"`
    Surname   string    `json:"surname"`
    Email     string    `json:"email"`
    DOB       time.Time `json:"dob"`
}