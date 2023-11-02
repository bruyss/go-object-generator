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
package generate

import (
	"os"
	"time"

	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/spf13/cobra"
)

// generateDigmonsCmd represents the digmons command
var generateDigmonsCmd = &cobra.Command{
	Use:   "digmons",
	Short: "Generate digmon objects.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now().Format("20060102_150405")
		obwriter.GenFolderName = obwriter.GenFolderRoot + "/" + now + "_digmons"
		err := os.MkdirAll(obwriter.GenFolderName, 0666)
		if err != nil && !os.IsExist(err) {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		digmonGen := obwriter.Generator{
			GeneralSettings: map[string]string{},
			ObjectSettings:  map[string]string{},
			Objects:         sheetreader.ReadDigmons(excelSource),
		}

		if genAll || genIdbs {
			idbTemplate, err := cmd.Flags().GetString("idb-template")
			if err != nil {
				logger.Sugar.Error(err)
			}
			idbFileName, err := cmd.Flags().GetString("idb-file")
			if err != nil {
				logger.Sugar.Error(err)
			}
			err = digmonGen.Generate(idbFileName, idbTemplate, tmp)
			if err != nil {
				logger.Sugar.Error(err)
			}
		}

		if genAll || genSource {
			sourceFileName, err := cmd.Flags().GetString("source-file")
			if err != nil {
				logger.Sugar.Error(err)
			}
			err = digmonGen.Generate(sourceFileName, "digmon.tmpl", tmp)
			if err != nil {
				logger.Sugar.Error(err)
			}
		}

		if genAll || genTags {
			logger.Sugar.Error("Gotta do those tag tables...")
		}

	},
}

func init() {
	GenerateCmd.AddCommand(generateDigmonsCmd)
}
