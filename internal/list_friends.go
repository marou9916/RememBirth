package internal

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// Fonction pour styliser les amis en tableau
func ListFriends() {
	friends := LoadFriends()
	if len(friends) == 0 {
		color.Red("Aucun ami trouvé.")
		return
	}

	// Créer le tableau
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Surname", "Birthday"})

	// Style des en-têtes (bleu-violet pour le fond, texte coloré)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgGreenColor},  // Fond bleu-violet avec texte vert pour "Nom"
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgYellowColor}, // Fond bleu-violet avec texte jaune pour "Prénom"
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiMagentaColor, tablewriter.FgCyanColor},   // Fond bleu-violet avec texte cyan pour "Anniversaire"
	)

	// Boucle pour ajouter chaque ami au tableau
	for _, friend := range friends {
		row := []string{
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Name),                          // Texte noir avec fond blanc
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Surname),                       // Texte noir avec fond blanc
			fmt.Sprintf("\033[30;47m%s\033[0m", friend.Birthday.Format("2006-01-02")), // Texte noir avec fond blanc
		}
		table.Append(row)
	}

	// Style du tableau
	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}
