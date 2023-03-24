package soa

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serialCmd represents the serial command
var serialCmd = &cobra.Command{
	Use:   "serial",
	Short: "get/set with serial in SOA of zone file",
	Long: `This subcommand to operate with serial value in SOA of zone file.
	For example:

	dnszone soa set serial 2023010101`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("serial called")
	},
}

func init() {

	serialCmd.AddCommand(setCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
