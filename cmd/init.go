/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the object generator",
	Long: `Command will create the following files to initialize object generation:
	- excelsource_go.xlsx: spreadsheet for entering object data
	- setting.json: settings for object generation
	- /templates: folder containing the object generation templates`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Are you sure you want to initialize [Y/N]: ")
		var answer string
		fmt.Scan(&answer)
		answer = strings.Replace(answer, "\n", "", -1)
		answer = strings.ToUpper(answer)
		for answer != "Y" && answer != "N" {
			fmt.Print(`Invalid input, please type "Y" or "N": `)
			fmt.Scan(&answer)
		}
		if answer == "Y" {
			return nil
		} else if answer == "N" {
			os.Exit(0)
		}
		return os.ErrInvalid
	},
	Run: func(cmd *cobra.Command, args []string) {
		flagExcel, _ := cmd.Flags().GetBool("excel")
		flagSettings, _ := cmd.Flags().GetBool("settings")
		flagTemplates, _ := cmd.Flags().GetBool("templates")
		if !(flagExcel || flagSettings || flagTemplates) {
			utils.Sugar.Info("Initializing all...")
			sheetreader.InitializeWorkbook(viper.GetString("filenames.general.objectsource"))
			viper.WriteConfig()
			obwriter.WriteTemplates(obwriter.Templates)
		} else {
			if flagExcel {
				utils.Sugar.Info("Initializing spreadsheet...")
				sheetreader.InitializeWorkbook(viper.GetString("filenames.general.objectsource"))
			}
			if flagSettings {
				utils.Sugar.Info("Initializing settings...")
				viper.WriteConfig()
			}
			if flagTemplates {
				utils.Sugar.Info("Initializing templates...")
				obwriter.WriteTemplates(obwriter.Templates)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initCmd.Flags().Bool("excel", false, "Initialize spreadsheet")
	initCmd.Flags().Bool("settings", false, "Initialize settings")
	initCmd.Flags().Bool("templates", false, "Initialize templates")

	initCmd.Flag("excel").NoOptDefVal = "true"
	initCmd.Flag("settings").NoOptDefVal = "true"
	initCmd.Flag("templates").NoOptDefVal = "true"

}
