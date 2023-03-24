package soa

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SoaCmd represents the soa command
var SoaCmd = &cobra.Command{
	Use:   "soa",
	Short: "Control SOA of zone file",
	Long: `This command controls SOA of zone file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("soa called")
	},
}

func init() {
	SoaCmd.AddCommand(serialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// soaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// soaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
