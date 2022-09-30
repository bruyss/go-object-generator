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
	"github.com/spf13/viper"
)

// generateAllCmd represents the generateAll command
var generateAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all objects",
	Run: func(cmd *cobra.Command, args []string) {

		// Define generators
		measmonGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.measmon"),
			Objects:         sheetreader.ReadMeasmons(excelSource),
		}
		digmonGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.digmon"),
			Objects:         sheetreader.ReadDigmons(excelSource),
		}
		valveGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.valve"),
			Objects:         sheetreader.ReadValves(excelSource),
		}
		controlValveGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.controlvalve"),
			Objects:         sheetreader.ReadControlValves(excelSource),
		}
		motorGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.motor"),
			Objects:         sheetreader.ReadMotors(excelSource),
		}
		freqMotorGen := obwriter.Generator{
			GeneralSettings: viper.GetStringMapString("gensettings.general"),
			ObjectSettings:  viper.GetStringMapString("gensettings.freqmotor"),
			Objects:         sheetreader.ReadFreqMotors(excelSource),
		}

		// Generate IDBs
		if genAll || genIdbs {
			utils.Sugar.Debugw("Generating IDBS",
				"genAll", genAll,
				"genIdbs", genIdbs)
			var err error
			idbTemplate := viper.GetString("filenames.general.idbtemplate")
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = measmonGen.Generate(viper.GetString("filenames.measmon.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = digmonGen.Generate(viper.GetString("filenames.digmon.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = valveGen.Generate(viper.GetString("filenames.valve.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = controlValveGen.Generate(viper.GetString("filenames.controlvalve.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = motorGen.Generate(viper.GetString("filenames.motor.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = freqMotorGen.Generate(viper.GetString("filenames.freqmotor.idbfile"), idbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		// Generate HMI DBs
		if genAll || genIdbs {
			var err error
			hmiDbTemplate := viper.GetString("filenames.general.hmidbtemplate")
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = measmonGen.Generate(viper.GetString("filenames.measmon.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = digmonGen.Generate(viper.GetString("filenames.digmon.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = valveGen.Generate(viper.GetString("filenames.valve.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = controlValveGen.Generate(viper.GetString("filenames.controlvalve.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = motorGen.Generate(viper.GetString("filenames.motor.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
			err = freqMotorGen.Generate(viper.GetString("filenames.freqmotor.hmidbfile"), hmiDbTemplate, tmp)
			if err != nil {
				utils.Sugar.Error(err)
			}
		}

		// Generate source files
		if genAll || genSource {
			utils.Sugar.Debugw("Generating source files",
				"genAll", genAll,
				"genSource", genSource)
			err := measmonGen.Generate(viper.GetString("filenames.measmon.sourcefile"), viper.GetString("filenames.measmon.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "measmons")
			}
			err = digmonGen.Generate(viper.GetString("filenames.digmon.sourcefile"), viper.GetString("filenames.digmon.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "digmons")
			}
			err = valveGen.Generate(viper.GetString("filenames.valve.sourcefile"), viper.GetString("filenames.valve.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "valves")
			}
			err = controlValveGen.Generate(viper.GetString("filenames.controlvalve.sourcefile"), viper.GetString("filenames.controlvalve.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "control valves")
			}
			err = motorGen.Generate(viper.GetString("filenames.motor.sourcefile"), viper.GetString("filenames.motor.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "motors")
			}
			err = freqMotorGen.Generate(viper.GetString("filenames.freqmotor.sourcefile"), viper.GetString("filenames.freqmotor.sourcetemplate"), tmp)
			if err != nil {
				utils.Sugar.Errorw(err.Error(), "generator", "frequency motors")
			}
		}

		// Generate tag tables
		if genAll || genTags {
			utils.Sugar.Debugw("Generating tag tables",
				"genAll", genAll,
				"genTags", genTags)
			var err error
			err = measmonGen.GeneratePlcTagTable(viper.GetString("filenames.measmon.tagfile"), "Measmons")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "measmons",
					"error", err)
			}
			err = digmonGen.GeneratePlcTagTable(viper.GetString("filenames.digmon.tagfile"), "Digmons")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "digmons",
					"error", err)
			}
			err = valveGen.GeneratePlcTagTable(viper.GetString("filenames.valve.tagfile"), "Valves")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "valves",
					"error", err)
			}
			err = controlValveGen.GeneratePlcTagTable(viper.GetString("filenames.controlvalve.tagfile"), "ControlValves")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "control valves",
					"error", err)
			}
			err = motorGen.GeneratePlcTagTable(viper.GetString("filenames.motor.tagfile"), "Motors")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "motors",
					"error", err)
			}
			err = freqMotorGen.GeneratePlcTagTable(viper.GetString("filenames.freqmotor.tagfile"), "FreqMotors")
			if err != nil {
				utils.Sugar.Errorw("Error generating tag table",
					"generator", "frequency motors",
					"error", err)
			}
		}

	},
}

func init() {
	generateCmd.AddCommand(generateAllCmd)
}
