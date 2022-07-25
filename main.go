package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/plc"
)

func main() {
	m := plc.Measmon{
		Tag:         "WWG-TT001",
		Description: "Test measmon",
		Unit:        "°C",
		Address:     "IW0",
		Direct:      false,
		LowLimit:    0,
		HighLimit:   150,
	}
	templateText, _ := ioutil.ReadFile("templates/measmon_template.txt")
	t := template.Must(template.New("measmon").Parse(string(templateText)))
	t.Execute(os.Stdout, m)
}
