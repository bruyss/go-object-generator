package obwriter

import (
	"log"
	"os"
)

const folderName = "templates"

func writeTemplate(name, template string) error {
	f, err := os.Create(folderName + "/" + name + ".tmpl")
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
	err := os.Mkdir(folderName, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	for k, v := range templates {
		err := writeTemplate(k, v)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
