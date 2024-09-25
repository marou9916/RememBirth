package internal

import (
	"fmt"
)

func ListFriends() {
	friends := LoadFriends()

	for _, friend := range friends {
		fmt.Printf("Nom: %s, Prénom: %s, Anniversaire: %s\n", friend.Nom, friend.Prenom, friend.DateAnniversaire.Format("2006-01-02"))
	}
}
