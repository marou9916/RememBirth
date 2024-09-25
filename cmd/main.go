package main

import (
	"fmt"
	"os"

	"github.com/marou9916/birthdayReminder.git/internal"
)

func main() {
	// Charger les amis depuis le fichier au début du programme
	friends := internal.LoadFriends()

	if len(os.Args) < 2 {
		fmt.Println("Usage: [add|list]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) != 5 {
			fmt.Println("Usage: add [name] [surname] [YYYY-MM-DD]")
			return
		}
		internal.AddFriend(os.Args[2], os.Args[3], os.Args[4], &friends) // Passer un pointeur
	case "list":
		internal.ListFriends()
	default:
		fmt.Println("Commande inconnue:", command)
	}
}
