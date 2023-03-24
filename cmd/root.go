package cmd

import (
	"os"

	"github.com/ip75/dnszone/commands"
	"github.com/ip75/dnszone/commands/record"
	"github.com/ip75/dnszone/commands/soa"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dnszone",
	Short: "CLI tool to edit DNS zone file",
	Long: `You can edit the DNS zone file
For example:
  1. change SOA serial key
    dnszone soa set serial 2023030301 
  2. add record to DNS zone
    dnszone record add "example.com A 1.1.1.1"
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(soa.SoaCmd, record.RecordCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnszone.yaml)")
	rootCmd.PersistentFlags().String(commands.GlobalFlagZoneFile, "zone-file.su", "specify DNS zone file to edit")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
