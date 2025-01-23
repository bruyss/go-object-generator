package obwriter

import (
	"encoding/xml"
	"io"
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/xuri/excelize/v2"
)

// GenFolderRoot is the default root directory for storing the generated files
const GenFolderRoot string = "genfiles"

// GenFolderName is the directory in the root directory where the generated files are stored
// At the start of generation, this value is filled with the current date & time
var GenFolderName string

// Generator represents the information needed to generate a certain object type.
//
// GeneralSettings contains settings that apply to all object types.
// ObjectSettings contains settings that apply to this specific object type.
// Objects is a slice of the objects to generate.
type Generator struct {
	GeneralSettings map[string]string
	ObjectSettings  map[string]string
	Objects         []plc.PlcObject
}

func (g *Generator) Generate(fileName, templateName string, tmp *template.Template) error {
	if len(g.Objects) == 0 {
		logger.Sugar.Debugw("No objects, not generating",
			"filename", fileName,
			"template", templateName,
		)
		return nil
	}

	f, err := os.Create(GenFolderRoot + "/" + fileName)
	if err != nil {
		return err
	}

	w := io.Writer(f)
	err = tmp.ExecuteTemplate(w, templateName, g)
	if err != nil {
		return err
	}

	logger.Sugar.Debugw("Generating",
		"filename", fileName,
		"template", templateName,
	)

	return nil
}

type xmlPlcTag struct {
	XMLName      xml.Name `xml:"Tagtable"`
	TagTableName string   `xml:"name,attr"`
	Tags         []xmlPlcTagLine
}

type xmlPlcTagLine struct {
	XMLName       xml.Name `xml:"Tag"`
	Tag           string   `xml:",innerxml"`
	Type          string   `xml:"type,attr"`
	HMIVisible    string   `xml:"hmiVisible,attr"`
	HMIWriteable  string   `xml:"hmiWriteable,attr"`
	HMIAccessible string   `xml:"hmiAccessible,attr"`
	Retain        string   `xml:"retain,attr"`
	Remark        string   `xml:"remark,attr"`
	Address       string   `xml:"addr,attr"`
}

func (g *Generator) GeneratePlcTagTable(fileName, tagTableName string) error {
	if len(g.Objects) == 0 {
		return nil
	}
	f, err := os.Create(GenFolderName + "/" + fileName)
	if err != nil {
		return err
	}
	tagtable := &xmlPlcTag{TagTableName: tagTableName}
	for _, o := range g.Objects {
		for _, t := range o.PlcTags() {
			tagtable.Tags = append(tagtable.Tags, xmlPlcTagLine{
				Type:          t.Dtype,
				Tag:           t.Name,
				HMIVisible:    "False",
				HMIWriteable:  "False",
				HMIAccessible: "False",
				Retain:        "False",
				Remark:        t.Comment,
				Address:       t.Address,
			})
		}
	}
	b, err := xml.MarshalIndent(&tagtable, "", "  ")
	if err != nil {
		return err
	}
	if _, err = f.Write(b); err != nil {
		return err
	}
	return nil
}

func setCellValue(f *excelize.File, sheet string, colIndex, rowIndex int, value interface{}) error {
	cellName, err := excelize.CoordinatesToCellName(colIndex, rowIndex)

	if err != nil {
		return err
	}

	return f.SetCellValue(sheet, cellName, value)
}

func (g *Generator) GenerateIoList(fileName string) error {
	f := excelize.NewFile()

	sheetName := f.GetSheetName(f.GetActiveSheetIndex())

	// Define columns for the IO list
	columns := []string{
		"Tag",
		"Description",
		"Address",
		"RIO",
		"Phase",
		"Status",
		"Comment",
	}

	f.SetSheetRow(sheetName, "A1", &columns)

	bottomRight, err := excelize.CoordinatesToCellName(len(columns), 2)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	enable := true
	err = f.AddTable(sheetName, &excelize.Table{
		Range:             "A1:" + bottomRight,
		Name:              "IO list",
		StyleName:         "TableStyleMedium2",
		ShowColumnStripes: false,
		ShowFirstColumn:   true,
		ShowHeaderRow:     &enable,
		ShowLastColumn:    false,
		ShowRowStripes:    &enable,
	})

	for _, object := range g.Objects {
		for colIndex, tag := range object.PlcTags() {
			setCellValue(f, sheetName, colIndex, 0, tag.Name)
			setCellValue(f, sheetName, colIndex, 1, tag.Comment)
			setCellValue(f, sheetName, colIndex, 2, tag.Address)
			setCellValue(f, sheetName, colIndex, 3, "")
			setCellValue(f, sheetName, colIndex, 4, "")
			setCellValue(f, sheetName, colIndex, 5, "To Test")
			setCellValue(f, sheetName, colIndex, 6, "")
		}
	}

	return f.SaveAs(fileName)
}
