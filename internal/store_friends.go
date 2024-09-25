package internal

import (
	"encoding/json"
	"os"

	"github.com/marou9916/birthdayReminder.git/models"
)

func StoreFriends(friends []models.Friend) {
	file, _ := json.MarshalIndent(friends, "", " ")
	_ = os.WriteFile("friends.json", file, 0644)
}
