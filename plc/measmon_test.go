package plc

import (
	"reflect"
	"testing"
)

func TestMeasmonInput(t *testing.T) {
	m := Measmon{
		Tag:         "WWG-TT001",
		Description: "Test measmon 1",
		Unit:        "°C",
		Address:     "IW0",
		Direct:      false,
		LowLimit:    0.0,
		HighLimit:   150.0,
	}
	want := map[string]string{
		"tag":         "WWG-TT001",
		"description": "Test measmon 1",
		"idb":         "IDB_WWG-TT001",
		"unit":        "°C",
		"input":       "WWG-TT001",
		"lowlim":      "0.0",
		"highlim":     "150.0",
	}
	if !reflect.DeepEqual(m.InputMap(), want) {
		t.Fatalf("m.InputMap() = %s, expected %s", m.InputMap(), want)
	}
}

func TestMeasmonTag(t *testing.T) {
	m := Measmon{
		Tag:         "WWG-TT001",
		Description: "Test measmon 1",
		Unit:        "°C",
		Address:     "IW0",
		Direct:      false,
		LowLimit:    0.0,
		HighLimit:   150.0,
	}
	want := []PlcTag{
		{"WWG-TT001", "Int", "IW0", "Test measmon 1"},
	}
	if !reflect.DeepEqual(m.PlcTags(), want) {
		t.Fatalf("m.PlcTags() = %s, expected %s", m.PlcTags(), want)
	}
}
