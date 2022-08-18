package plc

import (
	"reflect"
	"testing"
)

func TestNewDigmon(t *testing.T) {
	type args struct {
		tag         string
		description string
		address     string
		invert      string
		alarm       string
		invertAlarm string
	}
	tests := []struct {
		name    string
		args    args
		want    *digmon
		wantErr bool
	}{
		{
			"Digmon",
			args{"WWG-FS001", "Test flow switch", "I0.1", "false", "true", "true"},
			&digmon{"WWG-FS001", "Test flow switch", "I0.1", false, true, true},
			false,
		},
		{
			"Digmon no address",
			args{"WWG-FS001", "Test flow switch", "", "false", "true", "true"},
			&digmon{"WWG-FS001", "Test flow switch", "M0.0", false, true, true},
			false,
		},
		{
			"Digmon bad invert",
			args{"WWG-FS001", "Test flow switch", "", "allo", "true", "true"},
			nil,
			true,
		},
		{
			"Digmon bad alarm",
			args{"WWG-FS001", "Test flow switch", "", "false", "allo", "true"},
			nil,
			true,
		},
		{
			"Digmon bad invert alarm",
			args{"WWG-FS001", "Test flow switch", "", "false", "true", "allo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDigmon(tt.args.tag, tt.args.description, tt.args.address, tt.args.invert, tt.args.alarm, tt.args.invertAlarm)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDigmon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDigmon() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_digmon_Tag(t *testing.T) {
	tests := []struct {
		name string
		d    *digmon
		want string
	}{
		{
			"Digmon 1",
			&digmon{
				Tag:         "CTP-LSL001",
				Description: "Test digmon 1",
				Address:     "I3.0",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: false,
			},
			"CTP-LSL001",
		},
		{
			"Digmon 2",
			&digmon{
				Tag:         "CTP-LSL002",
				Description: "Test digmon 2",
				Address:     "",
				Invert:      true,
				Alarm:       true,
				InvertAlarm: false,
			},
			"CTP-LSL002",
		},
		{
			"Digmon 3",
			&digmon{
				Tag:         "CTP-LSL003",
				Description: "Test digmon 3",
				Address:     "I3.1",
				Invert:      false,
				Alarm:       false,
				InvertAlarm: false,
			},
			"CTP-LSL003",
		},
		{
			"Digmon 4",
			&digmon{
				Tag:         "CTP-LSH004",
				Description: "Test digmon 4",
				Address:     "",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: true,
			},
			"CTP-LSH004",
		},
		{
			"Digmon 5",
			&digmon{
				Tag:         "CTP-LSL005",
				Description: "Test digmon 5",
				Address:     "I3.2",
				Invert:      true,
				Alarm:       false,
				InvertAlarm: false,
			},
			"CTP-LSL005",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Digmon.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digmon_InputMap(t *testing.T) {
	tests := []struct {
		name string
		d    *digmon
		want map[string]string
	}{
		{
			"Digmon 1",
			&digmon{
				Tag:         "CTP-LSL001",
				Description: "Test digmon 1",
				Address:     "I3.0",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: false,
			},
			map[string]string{
				"Tag":         "CTP-LSL001",
				"Description": "Test digmon 1",
				"IDB":         "IDB_CTP-LSL001",
				"Input":       `"CTP-LSL001"`,
				"Invert":      "false",
				"Alarm":       "true",
				"InvertAlarm": "false",
			},
		},
		{
			"Digmon 2",
			&digmon{
				Tag:         "CTP-LSL002",
				Description: "Test digmon 2",
				Address:     "",
				Invert:      true,
				Alarm:       true,
				InvertAlarm: false,
			},
			map[string]string{
				"Tag":         "CTP-LSL002",
				"Description": "Test digmon 2",
				"IDB":         "IDB_CTP-LSL002",
				"Input":       `"CTP-LSL002"`,
				"Invert":      "true",
				"Alarm":       "true",
				"InvertAlarm": "false",
			},
		},
		{
			"Digmon 3",
			&digmon{
				Tag:         "CTP-LSL003",
				Description: "Test digmon 3",
				Address:     "I3.1",
				Invert:      false,
				Alarm:       false,
				InvertAlarm: false,
			},
			map[string]string{
				"Tag":         "CTP-LSL003",
				"Description": "Test digmon 3",
				"IDB":         "IDB_CTP-LSL003",
				"Input":       `"CTP-LSL003"`,
				"Invert":      "false",
				"Alarm":       "false",
				"InvertAlarm": "false",
			},
		},
		{
			"Digmon 4",
			&digmon{
				Tag:         "CTP-LSH004",
				Description: "Test digmon 4",
				Address:     "",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: true,
			},
			map[string]string{
				"Tag":         "CTP-LSH004",
				"Description": "Test digmon 4",
				"IDB":         "IDB_CTP-LSH004",
				"Input":       `"CTP-LSH004"`,
				"Invert":      "false",
				"Alarm":       "true",
				"InvertAlarm": "true",
			},
		},
		{
			"Digmon 5",
			&digmon{
				Tag:         "CTP-LSL005",
				Description: "Test digmon 5",
				Address:     "I3.2",
				Invert:      true,
				Alarm:       false,
				InvertAlarm: false,
			},
			map[string]string{
				"Tag":         "CTP-LSL005",
				"Description": "Test digmon 5",
				"IDB":         "IDB_CTP-LSL005",
				"Input":       `"CTP-LSL005"`,
				"Invert":      "true",
				"Alarm":       "false",
				"InvertAlarm": "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digmon.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digmon_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		d    *digmon
		want []*PlcTag
	}{
		{
			"Digmon 1",
			&digmon{
				Tag:         "CTP-LSL001",
				Description: "Test digmon 1",
				Address:     "I3.0",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL001", "Bool", "I3.0", "Test digmon 1"},
			},
		},
		{
			"Digmon 2",
			&digmon{
				Tag:         "CTP-LSL002",
				Description: "Test digmon 2",
				Address:     "I3.1",
				Invert:      true,
				Alarm:       true,
				InvertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL002", "Bool", "I3.1", "Test digmon 2"},
			},
		},
		{
			"Digmon 3",
			&digmon{
				Tag:         "CTP-LSL003",
				Description: "Test digmon 3",
				Address:     "I3.1",
				Invert:      false,
				Alarm:       false,
				InvertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL003", "Bool", "I3.1", "Test digmon 3"},
			},
		},
		{
			"Digmon 4",
			&digmon{
				Tag:         "CTP-LSH004",
				Description: "Test digmon 4",
				Address:     "I3.7",
				Invert:      false,
				Alarm:       true,
				InvertAlarm: true,
			},
			[]*PlcTag{
				{"CTP-LSH004", "Bool", "I3.7", "Test digmon 4"},
			},
		},
		{
			"Digmon 5",
			&digmon{
				Tag:         "CTP-LSL005",
				Description: "Test digmon 5",
				Address:     "I3.2",
				Invert:      true,
				Alarm:       false,
				InvertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL005", "Bool", "I3.2", "Test digmon 5"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.PlcTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digmon.PlcTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
