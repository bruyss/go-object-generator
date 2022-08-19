package obwriter

import (
	"io"
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/plc"
)

const genFolderName string = "genfiles"

type Generator struct {
	GeneralSettings generalSettings
	ObjectSettings  objectSettings
	Objects         []plc.PlcObject
}

func (g *Generator) Generate(fileName, templateName string, tmp *template.Template) error {
	err := os.Mkdir(genFolderName, 0666)
	if err != nil && !os.IsExist(err) {
		return err
	}
	f, err := os.Create(genFolderName + "/" + fileName)
	if err != nil {
		return err
	}

	w := io.Writer(f)
	err = tmp.ExecuteTemplate(w, templateName, g)
	if err != nil {
		return err
	}

	return nil
}
