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
	"os"
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/utils"
	"github.com/spf13/cobra"
)

var genSettings *obwriter.GeneratorSettings
var excelSource *excelize.File
var tmp *template.Template

var genAll, genIdbs, genSource, genTags bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Object generation",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		genSettings, err = obwriter.ReadSettings("settings.json")
		if err == os.ErrNotExist {
			obwriter.WriteDefaultSettings("settings.json")
			genSettings = obwriter.DefaultSettings
		} else if err != nil {
			return err
		}
		fileName, err := cmd.Flags().GetString("file")
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}
		excelSource, err = excelize.OpenFile(fileName)
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}
		tmp, err = template.ParseGlob("templates/*.tmpl")
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}

		genIdbs, err = cmd.Flags().GetBool("idbs")
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}
		genSource, err = cmd.Flags().GetBool("source-files")
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}
		genTags, err = cmd.Flags().GetBool("tag-tables")
		if err != nil {
			utils.Sugar.Error(err)
			return err
		}
		genAll = !genIdbs && !genSource && !genTags

		utils.Sugar.Debugw("Generate switch bits",
			"IDBs", genIdbs,
			"Source files", genSource,
			"Tags", genTags,
			"All", genAll)

		return nil
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("generate called")
	// },
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Persistent flags
	generateCmd.PersistentFlags().StringP("file", "f", "excelsource_go.xlsx", "File name of the spreadsheet containing object information")

	generateCmd.PersistentFlags().BoolP("idbs", "i", false, "Generate instance DBs.")
	generateCmd.PersistentFlags().BoolP("source-files", "s", false, "Generate source files.")
	generateCmd.PersistentFlags().BoolP("tag-tables", "t", false, "Generate tag tables.")

	generateCmd.Flag("idbs").NoOptDefVal = "true"
	generateCmd.Flag("source-files").NoOptDefVal = "true"
	generateCmd.Flag("tag-tables").NoOptDefVal = "true"

	generateCmd.PersistentFlags().String("idb-file", "", "Instance DB file name.")
	generateCmd.PersistentFlags().String("source-file", "", "Instance DB file name.")
	generateCmd.PersistentFlags().String("tag-file", "", "Instance DB file name.")
	generateCmd.PersistentFlags().String("idb-template", "idbs.tmpl", "Instance DB file template name.")
	generateCmd.PersistentFlags().String("source-template", "", "Source file template name.")
	generateCmd.PersistentFlags().String("tag-template", "tagTable.tmpl", "Tag table file template name.")

	generateCmd.Flag("idb-file").Hidden = true
	generateCmd.Flag("source-file").Hidden = true
	generateCmd.Flag("tag-file").Hidden = true

	generateCmd.Flag("idb-template").Hidden = true
	generateCmd.Flag("source-template").Hidden = true
	generateCmd.Flag("tag-template").Hidden = true

	// Local flags

}
