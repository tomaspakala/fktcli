package auth

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

// userNameCmd represents the user name command
var userNameCmd = &cobra.Command{
	Use:   "userName",
	Short: "userName generates random user name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		respTypeValueString := cmd.Flag(responseType.FlagName).Value.String()
		respType := responseType.ResponseType(respTypeValueString)
		fmt.Println(respType.ToType(gofakeit.Username()))
	},
}
