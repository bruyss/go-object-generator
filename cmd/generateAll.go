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

	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
	"github.com/spf13/cobra"
)

// generateAllCmd represents the generateAll command
var generateAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all objects",
	Run: func(cmd *cobra.Command, args []string) {

		// Define generators
		measmonGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.Measmon,
			Objects:         sheetreader.ReadMeasmons(excelSource),
		}

		// Generate IDBs
		if genAll || genIdbs {
			utils.Sugar.Debugw("Generating IDBS",
				"genAll", genAll,
				"genIdbs", genIdbs)
			idbTemplate, err := cmd.Flags().GetString("idb-template")
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = measmonGen.Generate("Measmon_IDBs.db", idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		// Generate source files
		if genAll || genSource {
			utils.Sugar.Debugw("Generating source files",
				"genAll", genAll,
				"genSource", genSource)
			err := measmonGen.Generate("Measmon_source.scl", "measmon.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		// Generate tag tables
		if genAll || genTags {
			utils.Sugar.Debugw("Generating tag tables",
				"genAll", genAll,
				"genTags", genTags)
			var err error
			if err != nil {
				utils.Sugar.Error(err)
			}
			fmt.Println("Gotta do those tag tables...")
		}

	},
}

func init() {
	generateCmd.AddCommand(generateAllCmd)

	// generateAllCmd.Flag("idbs").NoOptDefVal = "true"
	// generateAllCmd.Flag("source-files").NoOptDefVal = "true"
	// generateAllCmd.Flag("tag-tables").NoOptDefVal = "true"
}
