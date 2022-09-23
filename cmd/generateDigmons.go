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
	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
	"github.com/spf13/cobra"
)

// generateDigmonsCmd represents the digmons command
var generateDigmonsCmd = &cobra.Command{
	Use:   "digmons",
	Short: "Generate digmon objects.",
	Run: func(cmd *cobra.Command, args []string) {

		digmonGen := obwriter.Generator{
			GeneralSettings: map[string]string{},
			ObjectSettings:  map[string]string{},
			Objects:         sheetreader.ReadDigmons(excelSource),
		}

		if genAll || genIdbs {
			idbTemplate, err := cmd.Flags().GetString("idb-template")
			if err != nil {
				utils.Sugar.Error(err)
			}
			idbFileName, err := cmd.Flags().GetString("idb-file")
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = digmonGen.Generate(idbFileName, idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		if genAll || genSource {
			sourceFileName, err := cmd.Flags().GetString("source-file")
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = digmonGen.Generate(sourceFileName, "digmon.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		if genAll || genTags {
			utils.Logger.Info("Gotta do those tag tables...")
		}

	},
}

func init() {
	generateCmd.AddCommand(generateDigmonsCmd)

	generateDigmonsCmd.Flag("idb-file").Hidden = false
	generateDigmonsCmd.Flag("source-file").Hidden = false
	generateDigmonsCmd.Flag("tag-file").Hidden = false

	generateDigmonsCmd.Flag("idb-file").DefValue = "Digmon_IDBs.db"
	generateDigmonsCmd.Flag("source-file").DefValue = "Digmon_source.scl"
	generateDigmonsCmd.Flag("tag-file").DefValue = "Digmon_tags.xml"
}
