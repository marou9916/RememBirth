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
		fmt.Println("Usage: [ajouter|lister]")
		return
	}

	commande := os.Args[1]

	switch commande {
	case "ajouter":
		if len(os.Args) != 5 {
			fmt.Println("Usage: ajouter [nom] [prenom] [YYYY-MM-DD]")
			return
		}
		internal.AddFriend(os.Args[2], os.Args[3], os.Args[4], &friends) // Passer un pointeur
	case "lister":
		internal.ListFriends()
	default:
		fmt.Println("Commande inconnue:", commande)
	}
}
