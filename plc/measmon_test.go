package plc

import (
	"reflect"
	"testing"
)

func TestNewMeasmon(t *testing.T) {
	type args struct {
		tag         string
		description string
		unit        string
		address     string
		direct      string
		lowLimit    string
		highLimit   string
	}
	tests := []struct {
		name string
		args args
		want *measmon
	}{
		{
			"Measmon 1",
			args{tag: "WWG-TT001", description: "Test measmon 1", unit: "bar", address: "IW64", direct: "false", lowLimit: "-1.0", highLimit: "10.0"},
			&measmon{"WWG-TT001", "Test measmon 1", "bar", "IW64", false, -1.0, 10.0},
		},
		{
			"Measmon 2",
			args{tag: "WWG-TT002", description: "Test measmon 2", unit: "°C", address: "IW68", direct: "false", lowLimit: "-50.0", highLimit: "-100.0"},
			&measmon{"WWG-TT002", "Test measmon 2", "°C", "IW68", false, 0.0, 100.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMeasmon(tt.args.tag, tt.args.description, tt.args.unit, tt.args.address, tt.args.direct, tt.args.lowLimit, tt.args.highLimit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMeasmon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_measmon_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		m    *measmon
		want []*PlcTag
	}{
		{
			"Case 1",
			&measmon{
				Tag:         "WWG-TT001",
				Description: "Test measmon 1",
				Unit:        "°C",
				Address:     "IW16",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   100.0,
			},
			[]*PlcTag{
				{name: "WWG-TT001", dtype: "Int", address: "IW16", comment: "Test measmon 1"},
			},
		},
		{
			"Case 2",
			&measmon{
				Tag:         "WWG-FT656",
				Description: "Test measmon 2",
				Unit:        "°C",
				Address:     "IW18",
				Direct:      false,
				LowLimit:    0.0,
				HighLimit:   150.0,
			},
			[]*PlcTag{
				{name: "WWG-FT656", dtype: "Int", address: "IW18", comment: "Test measmon 2"},
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

func Test_measmon_InputMap(t *testing.T) {
	tests := []struct {
		name string
		m    *measmon
		want map[string]string
	}{
		{
			"Case 1",
			&measmon{
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
				"Input":       `"WWG-TT001"`,
				"LowLimit":    "0.0",
				"HighLimit":   "100.0",
			},
		},
		{
			"Case 2",
			&measmon{
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
				"Input":       `"WWG-FT656"`,
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
