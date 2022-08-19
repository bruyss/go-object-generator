package main

import (
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
)

func init() {
	utils.InitializeCustomLogger()
}

func main() {
	// sheetreader.InitializeWorkbook("excelsource_go.xlsx")
	f, _ := excelize.OpenFile("excelsource_go.xlsx")

	// obwriter.WriteDefaultSettings("settings.json")
	// obwriter.WriteTemplates(obwriter.Templates)
	settings := obwriter.ReadSettings("settings.json")

	measmonGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.Measmon,
		Objects:         sheetreader.ReadMeasmons(f),
	}

	digmonGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.Digmon,
		Objects:         sheetreader.ReadDigmons(f),
	}
	valveGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.Valve,
		Objects:         sheetreader.ReadValves(f),
	}
	controlValveGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.ControlValve,
		Objects:         sheetreader.ReadControlValves(f),
	}
	motorGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.Motor,
		Objects:         sheetreader.ReadMotors(f),
	}
	freMotorGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.FreqMotor,
		Objects:         sheetreader.ReadFreqMotors(f),
	}

	tmp, err := template.ParseGlob("templates/*")

	if err != nil {
		utils.Sugar.Error(err)
	}

	// err = tmp.Execute(os.Stdout, measGen)

	// if err != nil {
	// 	utils.Sugar.Error(err)
	// }

	measmonGen.Generate("Measmon_IDBs.db", "idbs.tmpl", tmp)
	measmonGen.Generate("Measmon_source.scl", "measmon.tmpl", tmp)
	digmonGen.Generate("Digmon_IDBs.db", "idbs.tmpl", tmp)
	valveGen.Generate("Valve_IDBs.db", "idbs.tmpl", tmp)
	controlValveGen.Generate("ControlValve_IDBs.db", "idbs.tmpl", tmp)
	motorGen.Generate("Motor_IDBs.db", "idbs.tmpl", tmp)
	freMotorGen.Generate("FreqMotor_IDBs.db", "idbs.tmpl", tmp)

}
