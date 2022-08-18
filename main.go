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
	measmons := sheetreader.ReadMeasmons(f)

	obwriter.WriteDefaultSettings("settings.json")
	obwriter.WriteTemplates(obwriter.Templates)
	settings := obwriter.ReadSettings("settings.json")

	measGen := obwriter.Generator{
		GeneralSettings: *settings.General,
		ObjectSettings:  *settings.Measmon,
		Objects:         measmons,
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

}
