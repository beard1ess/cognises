package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateHostsCmd represents the generateHosts command
var generateHostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Generate a Linux Hosts table from cloud servers.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented..")
	},
}

func init() {
	// NYI
	//rootCmd.AddCommand(generateHostsCmd)
}
