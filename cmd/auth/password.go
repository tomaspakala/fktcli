package auth

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

var (
	lower   bool
	upper   bool
	numeric bool
	special bool
	length  int32
)

// passwordCmd represents the user name command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "password generates random user name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := gofakeit.Password(lower, upper, numeric, special, false, int(length))
		fmt.Println(responseType.GetValueByResponseType(cmd.Flags(), result))
	},
}

func init() {
	passwordCmd.Flags().BoolVarP(&lower, "lower", "l", true, "include lower")
	passwordCmd.Flags().BoolVarP(&upper, "upper", "u", true, "include upper")
	passwordCmd.Flags().BoolVarP(&numeric, "numeric", "n", true, "include numeric")
	passwordCmd.Flags().BoolVarP(&special, "special", "s", true, "include special")
	passwordCmd.Flags().Int32VarP(&length, "length", "c", 14, "password length")
}
