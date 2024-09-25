package internal

import (
	"encoding/json"
	"os"

	"github.com/marou9916/birthdayReminder.git/models"
)

func StoreFriends(amis []models.Friend) {
	file, _ := json.MarshalIndent(amis, "", " ")
	_ = os.WriteFile("amis.json", file, 0644)
}
