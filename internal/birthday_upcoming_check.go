package internal

import (
	"time"

	"github.com/marou9916/birthdayReminder.git/models"
)

// CheckUpcomingBirthdays retourne une liste d'amis avec un anniversaire à venir
func CheckUpcomingBirthdays(friends []models.Friend, days int) []models.Friend {
	var upcoming []models.Friend
	today := time.Now()
	limitDate := today.AddDate(0, 0, days) // Ajouter le nombre de jours spécifié

	for _, friend := range friends {
		// On compare uniquement le jour et le mois
		if friend.Birthday.Month() == limitDate.Month() && friend.Birthday.Day() == limitDate.Day() {
			upcoming = append(upcoming, friend)
		}
	}

	return upcoming
}
