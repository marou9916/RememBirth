package internal

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/marou9916/birthdayReminder.git/models"
	"github.com/olekukonko/tablewriter"
)

// Function to list friends in a styled table
func ListFriends(friends []models.Friend) {
	if len(friends) == 0 {
		color.Red("No friends found.")
		return
	}

	// Create the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Surname", "Birthday"}) // Removed Email header

	// Header style (blue-violet background, colored text)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgGreenColor},  // Background blue-violet with green text for "Name"
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgYellowColor}, // Background blue-violet with yellow text for "Surname"
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgCyanColor},   // Background blue-violet with cyan text for "Birthday"
	)

	// Loop to add each friend to the table
	for _, friend := range friends {
		row := []string{
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Name),                          // Black text with white background
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Surname),                       // Black text with white background
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Birthday.Format("2006-01-02")), // Black text with white background
		}
		table.Append(row)
	}

	// Table style
	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}
