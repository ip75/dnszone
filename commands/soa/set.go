package soa

import (
	"fmt"
	"strconv"

	"github.com/ip75/dnszone/commands"
	"github.com/shuLhan/share/lib/dns"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set serial in SOA of DNS zone file",
	Long: `This command sets serial number in SOA of DNS zone file:
For example:

 dnszone soa set serial 2023010101`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("no serial parameter of SOA specified", cmd.Usage())
			return
		}

		serial, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			fmt.Printf("unable to parse serial argument: %v\nIt must me integer", err)
			return
		}

		zoneFile := cmd.Flag(commands.GlobalFlagZoneFile)

		if !zoneFile.Changed {
			fmt.Println(zoneFile.Usage)
			return
		}

		zoneFileName := zoneFile.Value.String()

		zone, err := dns.ParseZoneFile(zoneFileName, "bvgm.su.", 3600)
		if err != nil {
			fmt.Printf("parse zone file error: %v", err)
			return
		}

		zone.SOA.Serial = uint32(serial)

		err = zone.Save()
		if err != nil {
			fmt.Printf("save zone file error: %v", err)
			return
		}

		fmt.Println("serial of SOA is set succesfully")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
