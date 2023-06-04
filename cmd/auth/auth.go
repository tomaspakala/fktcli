package auth

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AuthCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth generates random data related to auth",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(cmd.Help())
	},
}

func addSubcommandPalettes() {
	AuthCmd.AddCommand(userNameCmd)
	AuthCmd.AddCommand(passwordCmd)
}

func init() {
	addSubcommandPalettes()
}
