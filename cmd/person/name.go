package person

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "name generates random name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := gofakeit.Name()
		fmt.Println(responseType.GetValueByResponseType(cmd.Flags(), result))
	},
}
