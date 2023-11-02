package obwriter

import (
	"encoding/xml"
	"io"
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/plc"
)

const GenFolderRoot string = "genfiles"

var GenFolderName string

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
