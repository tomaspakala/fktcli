package cmd

import (
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
	"github.com/tomaspakala/fktcli/cmd/auth"
	"github.com/tomaspakala/fktcli/cmd/datetime"
	"github.com/tomaspakala/fktcli/cmd/guid"
	"github.com/tomaspakala/fktcli/cmd/person"
	responseType "github.com/tomaspakala/fktcli/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fktcli",
	Short: "fktcli generates various fake data",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(guid.GuidCmd)
	rootCmd.AddCommand(datetime.DateTimeCmd)
	rootCmd.AddCommand(person.PersonCmd)
	rootCmd.AddCommand(auth.AuthCmd)
}

func init() {
	gofakeit.Seed(time.Now().UnixNano())
	addSubcommandPalettes()

	var respType = responseType.NormalCase
	rootCmd.PersistentFlags().VarP(&respType, responseType.FlagName, "r",
		`Response types. Allowed: "normalCase", "upperCase", "lowerCase"`)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
