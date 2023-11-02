package obwriter

import (
	"os"

	"github.com/bruyss/go-object-generator/logger"
)

// templateFolderName contains the default name for the folder containing generation templates
const templateFolderName = "templates"

// writeTemplate writes a template string to a .tmpl file with a given name
func writeTemplate(name, template string) error {
	f, err := os.Create(templateFolderName + "/" + name + ".tmpl")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(template)
	if err != nil {
		return err
	}
	return nil
}

// WriteTemplates writes a list of templates to files
// If the default storage folder does not exist it is created
func WriteTemplates(templates map[string]string) error {
	err := os.Mkdir(templateFolderName, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	for k, v := range templates {
		err := writeTemplate(k, v)
		if err != nil {
			logger.Sugar.Error(err)
		}
	}
	return nil
}
