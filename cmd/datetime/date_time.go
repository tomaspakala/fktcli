package datetime

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	responseType "github.com/tomaspakala/fktcli/util"
)

var timeNow = time.Now()

// DateTimeCmd represents the date time command
var DateTimeCmd = &cobra.Command{
	Use:   "dateTime",
	Short: "dateTime generating command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dateTime := getDateTime()
		fmt.Println(responseType.GetValueByResponseType(cmd.Flags(), dateTime))
	},
}

func getDateTime() string {
	return timeNow.UTC().Format(time.RFC3339Nano)
}
