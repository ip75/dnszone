package record

import (
	"fmt"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var RecordCmd = &cobra.Command{
	Use:   "record",
	Short: "Control records in DNS zone file",
	Long: `You can control records in DNS zone file. For example:

  add record:
    dnszone record add "example.com A 1.1.1.1"

  delete record:
    dnszone record add "example.com A 1.1.1.1"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("record called")
	},
}

func init() {
	RecordCmd.AddCommand(deleteCmd, addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
