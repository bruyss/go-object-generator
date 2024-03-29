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
*/package generate

import (
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
)

var (
	excelSource *excelize.File
	tmp         *template.Template
)

var genAll, genIdbs, genSource, genTags, genHMI bool

// GenerateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate PLC object source files based on data read from spreadsheet",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logger.Sugar.Info("")
		logger.Sugar.Info("")
		logger.Sugar.Info("Starting go-object-generator")

		var err error
		fileName := viper.GetString("filenames.general.objectsource")
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}
		excelSource, err = excelize.OpenFile(fileName)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}
		tmp, err = template.ParseGlob("templates/*.tmpl")
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}

		genAll = !genIdbs && !genSource && !genTags

		obwriter.GenFolderName = obwriter.GenFolderRoot
		err = os.MkdirAll(obwriter.GenFolderName, 0666)
		if err != nil && !os.IsExist(err) {
			return err
		}

		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		excelSource.Close()
		logger.Sugar.Info("Generation complete")
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("generate called")
	// },
}

func init() {
	// Persistent flags
	GenerateCmd.PersistentFlags().BoolVarP(&genIdbs, "idbs", "i", false, "Generate instance DBs.")
	GenerateCmd.PersistentFlags().BoolVarP(&genHMI, "hmiDBs", "d", false, "Generate HMI DB.")
	GenerateCmd.PersistentFlags().
		BoolVarP(&genSource, "source-files", "s", false, "Generate source files.")
	GenerateCmd.PersistentFlags().
		BoolVarP(&genTags, "tag-tables", "t", false, "Generate tag tables.")

	GenerateCmd.Flag("idbs").NoOptDefVal = "true"
	GenerateCmd.Flag("hmiDBs").NoOptDefVal = "true"
	GenerateCmd.Flag("source-files").NoOptDefVal = "true"
	GenerateCmd.Flag("tag-tables").NoOptDefVal = "true"
}
