package person

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "email generates random email",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := gofakeit.Email()
		fmt.Println(responseType.GetValueByResponseType(cmd.Flags(), result))
	},
}
