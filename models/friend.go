package models

import "time"

type Friend struct {
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Birthday          time.Time `json:"birthday"`
}
