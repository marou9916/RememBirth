package main

import (
	"fmt"
	"os"

	"github.com/marou9916/birthdayReminder.git/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [add|list|remind|delete]")
		return
	}

	command := os.Args[1]

	// Load friends from the file
	friends := internal.LoadFriends()

	switch command {
	case "add":
		if len(os.Args) != 5 {
			fmt.Println("Usage: add [name] [surname] [YYYY-MM-DD]")
			return
		}
		internal.AddFriend(os.Args[2], os.Args[3], os.Args[4], &friends)
	case "list":
		internal.ListFriends(friends)
	case "remind":
		upcoming := internal.CheckUpcomingBirthdays(friends, 1) //1 days in advance
		for _, friend := range upcoming {
			internal.SendBirthdayReminderEmail(friend) // Send the birthday reminder
		}
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: delete [name]")
			return
		}
		internal.DeleteFriend(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
	}
}
