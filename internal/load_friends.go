package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marou9916/birthdayReminder.git/models"
)

func LoadFriends() []models.Friend {
	var friends []models.Friend

	file, err := os.ReadFile("friends.json")
	if err != nil {
		fmt.Println("Aucun fichier trouvé, création d'une nouvelle liste.")
		return friends
	}

	err = json.Unmarshal(file, &friends)
	if err != nil {
		fmt.Println("Erreur de chargement des amis:", err)
		return friends
	}

	return friends
}


















