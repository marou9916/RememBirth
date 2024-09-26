package models

import "time"

// Friend represents a friend with their name, surname, birthday, and email address
type Friend struct {
    Name             string    `json:"name"`             // Updated to "name"
    Surname          string    `json:"surname"`          // Updated to "surname"
    Birthday         time.Time `json:"birthday"`         // Updated to "birthday"
    Email            string    `json:"email"`            // Field for the email address
}
