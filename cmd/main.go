package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marou9916/birthdayReminder.git/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [add|list|remind]")
		return
	}

	command := os.Args[1]

	// Load friends from the file
	friends := internal.LoadFriends()

	switch command {
	case "add":
		if len(os.Args) != 6 {
			fmt.Println("Usage: add [name] [surname] [YYYY-MM-DD] [email]")
			return
		}
		internal.AddFriend(os.Args[2], os.Args[3], os.Args[4], os.Args[5], &friends)

	case "list":
		internal.ListFriends(friends) // Call the ListFriends function to display friends

	case "remind":
		if len(os.Args) != 3 {
			fmt.Println("Usage: remind [days]")
			return
		}

		days, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Please provide a valid number of days.")
			return
		}

		upcoming := internal.CheckUpcomingBirthdays(friends, days) // Check upcoming birthdays
		for _, friend := range upcoming {
			internal.SendBirthdayReminder(friend) // Send the birthday reminder
		}
	
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: delete [surname]")
			return
		}
		internal.DeleteFriend(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}
