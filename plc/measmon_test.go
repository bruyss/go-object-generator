package plc

import (
	"reflect"
	"testing"
)

func TestMeasmon_Stringer(t *testing.T) {
	tests := []struct {
		name string
		m    *Measmon
		want string
	}{
		{
			"Case 1",
			&Measmon{
				Tag:         "WWG-TT001",
				Description: "Test measmon 1",
				Unit:        "°C",
				Address:     "IW16",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   100.0,
			},
			"WWG-TT001",
		},
		{
			"Case 2",
			&Measmon{
				Tag:         "WWG-FT656",
				Description: "Test measmon 2",
				Unit:        "°C",
				Address:     "IW18",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   150.0,
			},
			"WWG-FT656",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Stringer(); got != tt.want {
				t.Errorf("Measmon.Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeasmon_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		m    *Measmon
		want []PlcTag
	}{
		{
			"Case 1",
			&Measmon{
				Tag:         "WWG-TT001",
				Description: "Test measmon 1",
				Unit:        "°C",
				Address:     "IW16",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   100.0,
			},
			[]PlcTag{
				{
					name:    "WWG-TT001",
					dtype:   "Int",
					address: "IW16",
					comment: "Test measmon 1",
				},
			},
		},
		{
			"Case 2",
			&Measmon{
				Tag:         "WWG-FT656",
				Description: "Test measmon 2",
				Unit:        "°C",
				Address:     "IW18",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   150.0,
			},
			[]PlcTag{
				{
					name:    "WWG-FT656",
					dtype:   "Int",
					address: "IW18",
					comment: "Test measmon 2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.PlcTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Measmon.PlcTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeasmon_InputMap(t *testing.T) {
	tests := []struct {
		name string
		m    *Measmon
		want map[string]string
	}{
		{
			"Case 1",
			&Measmon{
				Tag:         "WWG-TT001",
				Description: "Test measmon 1",
				Unit:        "°C",
				Address:     "IW16",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   100.0,
			},
			map[string]string{
				"Tag":         "WWG-TT001",
				"Description": "Test measmon 1",
				"IDB":         "IDB_WWG-TT001",
				"Unit":        "°C",
				"Input":       "WWG-TT001",
				"LowLimit":    "0.0",
				"HighLimit":   "100.0",
			},
		},
		{
			"Case 2",
			&Measmon{
				Tag:         "WWG-FT656",
				Description: "Test measmon 2",
				Unit:        "m³/h",
				Address:     "IW18",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   150.0,
			},
			map[string]string{
				"Tag":         "WWG-FT656",
				"Description": "Test measmon 2",
				"IDB":         "IDB_WWG-FT656",
				"Unit":        "m³/h",
				"Input":       "WWG-FT656",
				"LowLimit":    "0.0",
				"HighLimit":   "150.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Measmon.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
