package person

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PersonCmd represents the person command
var PersonCmd = &cobra.Command{
	Use:   "person",
	Short: "person generates random data related to person",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(cmd.Help())
	},
}

func addSubcommandPalettes() {
	PersonCmd.AddCommand(nameCmd)
	PersonCmd.AddCommand(emailCmd)
}

func init() {
	addSubcommandPalettes()
}
