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

// func init() {
// 	_ = os.Mkdir(genFolderName, 0750)
// }

func (g *Generator) GenerateIDBs(fileName string, tmp *template.Template) error {
	f, err := os.Create(genFolderName + "/" + fileName + ".db")
	if err != nil {
		return err
	}

	w := io.Writer(f)
	err = tmp.Execute(w, g)
	if err != nil {
		return err
	}

	return nil
}
