package internal

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func ListFriends() {
	friends := LoadFriends()

	if len(friends) == 0 {
		color.Red("Aucun ami trouvé.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Surname", "Birthday"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
	)

	for _, friend := range friends {
		row := []string{
			friend.Name,
			friend.Surname,
			friend.Birthday.Format("2006-01-02"),
		}
		table.Append(row)
	}

	table.SetBorder(true)
	table.SetRowLine(true)

	table.Render()
}
