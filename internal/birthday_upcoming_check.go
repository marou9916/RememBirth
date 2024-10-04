package internal

import (
	"time"

	"github.com/marou9916/birthdayReminder.git/models"
)

// CheckUpcomingBirthdays returns a list of friends with a birthday in the next "days" days (e.g., 1 day for tomorrow)
func CheckUpcomingBirthdays(friends []models.Friend, days int) []models.Friend {
	var upcoming []models.Friend
	today := time.Now()
	limitDate := today.AddDate(0, 0, days) // Add the specified number of days

	for _, friend := range friends {
		// Create the friend's birthday for this year
		birthdayOfMyFriendThisYear := time.Date(today.Year(), friend.Birthday.Month(), friend.Birthday.Day(), 0, 0, 0, 0, time.UTC)

		// Check if the birthday is within the next "days" days
		if birthdayOfMyFriendThisYear.After(today) && birthdayOfMyFriendThisYear.Before(limitDate) {
			upcoming = append(upcoming, friend)
		}
	}

	return upcoming
}
