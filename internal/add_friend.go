package internal

import (
	"fmt"
	"time"

	"github.com/marou9916/birthdayReminder.git/models"
)

func AddFriend(nom, prenom, dateStr string, friends *[]models.Friend) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Erreur de format de date. Utilisez YYYY-MM-DD.")
		return
	}

	friend := models.Friend{
		Nom:              nom,
		Prenom:           prenom,
		DateAnniversaire: date,
	}

	*friends = append(*friends, friend)
	StoreFriends(*friends)

	fmt.Println("Ami ajouté avec succès!")
}
