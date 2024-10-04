package internal

import (
	"fmt"
	"time"

	"github.com/marou9916/birthdayReminder.git/models"
)

// Function to add a friend
func AddFriend(name, surname, dateStr string, friends *[]models.Friend) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Error in date format. Use YYYY-MM-DD.")
		return
	}

	friend := models.Friend{
		Name:     name,
		Surname:  surname,
		Birthday: date,
	}

	*friends = append(*friends, friend)
	StoreFriends(*friends) // Store the updated list of friends

	fmt.Println("Friend added successfully!")
}
