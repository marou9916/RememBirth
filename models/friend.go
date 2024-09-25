package models

import "time"

type Friend struct {
	Nom              string    `json:"nom"`
	Prenom           string    `json:"prenom"`
	DateAnniversaire time.Time `json:"date_anniversaire"`
}
