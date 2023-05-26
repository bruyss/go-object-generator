/*
Copyright © 2023 Jeroen Van Bruyssel <jeroen.vb1@gmail.com>
*/

package cmd

import (
	"os"

	"github.com/bruyss/go-object-generator/cmd/generate"
	"github.com/bruyss/go-object-generator/cmd/initialize"
	"github.com/bruyss/go-object-generator/config"
	"github.com/bruyss/go-object-generator/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var Version string = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "go-object-generator",
	Version: Version,
	Short:   "PLC object generator",
	Long: `A generator for PLC objects, reads data from a spreadsheet and outputs source files.

To initialize a new object generator project run: go-object-generator.exe init

To generate all objects entered in the spreadsheet run: go-object-generator.exe generate all

For further information the the available commands run: go-object-generator.exe --help`,
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
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(initialize.InitCmd)
	rootCmd.AddCommand(generate.GenerateCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-object-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Set default configuration
	config.SetDefaults()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".go-object-generator" (without extension).
		// viper.AddConfigPath(".")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		if err := viper.SafeWriteConfig(); err != nil {
			logger.Sugar.Warn(err)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		logger.Sugar.Info("Config file found")
	}
}
