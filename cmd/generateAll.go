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
		digmonGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.Digmon,
			Objects:         sheetreader.ReadDigmons(excelSource),
		}
		valveGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.Valve,
			Objects:         sheetreader.ReadValves(excelSource),
		}
		controlValveGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.ControlValve,
			Objects:         sheetreader.ReadControlValves(excelSource),
		}
		motorGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.Motor,
			Objects:         sheetreader.ReadMotors(excelSource),
		}
		freqMotorGen := obwriter.Generator{
			GeneralSettings: *genSettings.General,
			ObjectSettings:  *genSettings.FreqMotor,
			Objects:         sheetreader.ReadFreqMotors(excelSource),
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
			err = digmonGen.Generate("Digmon_IDBs.db", idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = valveGen.Generate("Valve_IDBs.db", idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = controlValveGen.Generate("ControlValve_IDBs.db", idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = motorGen.Generate("Motor_IDBs.db", idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = freqMotorGen.Generate("FreqMotor_IDBs.db", idbTemplate, tmp)
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
			err = digmonGen.Generate("Digmon_source.scl", "digmon.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = valveGen.Generate("Valve_source.scl", "valve.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = controlValveGen.Generate("ControlValve_source.scl", "controlValve.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = motorGen.Generate("Motor_source.scl", "motor.tmpl", tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = freqMotorGen.Generate("FreqMotor_source.scl", "freqMotor.tmpl", tmp)
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
			measmonGen.GeneratePlcTagTable("Measmon_tags.xml", "Measmons")
			digmonGen.GeneratePlcTagTable("Digmon_tags.xml", "Digmons")
			valveGen.GeneratePlcTagTable("Valve_tags.xml", "Valves")
			controlValveGen.GeneratePlcTagTable("ControlValve_tags.xml", "ControlValves")
			motorGen.GeneratePlcTagTable("Motor_tags.xml", "Motors")
			freqMotorGen.GeneratePlcTagTable("FreqMotor_tags.xml", "FreqMotors")
		}

	},
}

func init() {
	generateCmd.AddCommand(generateAllCmd)

	// generateAllCmd.Flag("idbs").NoOptDefVal = "true"
	// generateAllCmd.Flag("source-files").NoOptDefVal = "true"
	// generateAllCmd.Flag("tag-tables").NoOptDefVal = "true"
}