package guid

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

// GuidCmd represents the guid command
var GuidCmd = &cobra.Command{
	Use:   "guid",
	Short: "guid generating command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		guid := getGuid()
		fmt.Println(responseType.GetValueByResponseType(cmd.Flags(), guid))
	},
}

func getGuid() string {
	return uuid.New().String()
}
