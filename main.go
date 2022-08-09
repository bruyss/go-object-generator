package main

import (
	"log"
	"os"
	"text/template"

	"github.com/bruyss/go-object-generator/plc"
)

func main() {
	m := &plc.Measmon{
		Tag:         "WWG-TT001",
		Description: "Test measmon",
		Unit:        "°C",
		Address:     "IW0",
		Direct:      false,
		LowLimit:    0,
		HighLimit:   150,
	}
	m2 := &plc.Measmon{
		Tag:         "WWG-TT002",
		Description: "Test measmon",
		Unit:        "°C",
		Address:     "IW2",
		Direct:      false,
		LowLimit:    0,
		HighLimit:   100,
	}
	funcMap := template.FuncMap{
		"InputMap": plc.PLCObject.InputMap,
	}
	// t, err := template.ParseGlob("templates/*.tmpl")
	t, err := template.New("measmon.tmpl").Funcs(funcMap).ParseFiles("templates/measmon.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	measmons := []plc.PLCObject{m, m2}
	// measmons := []map[string]string{m.InputMap(), m2.InputMap()}
	t.ExecuteTemplate(os.Stdout, "measmon.tmpl", measmons)
}
