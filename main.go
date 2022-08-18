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

	obwriter.WriteDefaultSettings("settings.json")
	obwriter.WriteTemplates(obwriter.Templates)
	settings := obwriter.ReadSettings("settings.json")

	measGen := obwriter.Generator{
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

	tmp, err := template.ParseFiles("templates/idbs.tmpl")

	if err != nil {
		utils.Sugar.Error(err)
	}

	// err = tmp.Execute(os.Stdout, measGen)

	// if err != nil {
	// 	utils.Sugar.Error(err)
	// }

	measGen.GenerateIDBs("Measmon_IDBs", tmp)
	digmonGen.GenerateIDBs("Digmon_IDBs", tmp)
	valveGen.GenerateIDBs("Valve_IDBs", tmp)
	controlValveGen.GenerateIDBs("ControlValve_IDBs", tmp)
	motorGen.GenerateIDBs("Motor_IDBs", tmp)
	freMotorGen.GenerateIDBs("FreqMotor_IDBs", tmp)

}
