package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marou9916/birthdayReminder.git/models"
)

func LoadFriends() []models.Friend {
	var friends []models.Friend

	file, err := os.ReadFile("amis.json")
	if err != nil {
		// Si le fichier n'existe pas, retourner une liste vide
		fmt.Println("Aucun fichier trouvé, création d'une nouvelle liste.")
		return friends
	}

	// Charger les amis depuis le fichier JSON
	err = json.Unmarshal(file, &friends)
	if err != nil {
		fmt.Println("Erreur de chargement des amis:", err)
		return friends
	}

	return friends
}
