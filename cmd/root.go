package cmd

import (
	"fmt"
	"os"

	config "../config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var providers []string
var log = logrus.New()

var v = viper.New()
var c = config.CognisesConfiguration{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cognises",
	Short: "SSH Config generator and more.",

	//Run: func(cognises *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cognises.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		// Find home directory.

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cognises" (without extension).
		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName(".cognises")
	}

	v.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		err := v.Unmarshal(&c)
		if err != nil {
			log.Fatalf("Unable to load config, please check syntax. \n%s", err)
		}
	} else {
		c = config.CognisesConfiguration(config.DefaultConfig)
	}
}
