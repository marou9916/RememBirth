package models

import "time"

// Friend represents a friend with their name, surname, and birthday
type Friend struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Birthday time.Time `json:"birthday"`
}
