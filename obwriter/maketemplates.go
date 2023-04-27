package obwriter

import (
	"os"

	"github.com/bruyss/go-object-generator/logger"
)

const templateFolderName = "templates"

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
