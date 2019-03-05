package cmd

import (
	"fmt"

	config "../config"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// generateConfigCmd represents the generateConfig command
var generateConfigCmd = &cobra.Command{
	Use:   "generate-config",
	Short: "Writes default config to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := yaml.Marshal(&config.DefaultConfig)
		check(err)
		fmt.Println(string(d))
	},
}

func init() {
	rootCmd.AddCommand(generateConfigCmd)
}
