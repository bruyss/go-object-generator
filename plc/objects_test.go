package plc

import (
	"reflect"
	"testing"
)

func TestDigmonInput(t *testing.T) {
	d := Digmon{
		tag:         "WWG-LS001",
		description: "Test digmon 1",
		address:     "I0.0",
		invert:      false,
		alarm:       true,
		invertAlarm: false,
	}
	mapSuccess := map[string]string{
		"tag":         "WWG-TT001",
		"description": "Test digmon 1",
		"idb":         "IDB_WWG-LS001",
		"input":       "WWG-LS001",
	}
	if !reflect.DeepEqual(d.InputMap(), mapSuccess) {
		t.Fatalf("m.InputMap() = %s, expected %s", d.InputMap(), mapSuccess)
	}
}
