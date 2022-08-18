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
		invert      bool
		alarm       bool
		invertAlarm bool
	}
	tests := []struct {
		name    string
		args    args
		want    *digmon
		wantErr bool
	}{
		{
			"Digmon",
			args{"WWG-FS001", "Test flow switch", "I0.1", false, true, true},
			&digmon{"WWG-FS001", "Test flow switch", "I0.1", false, true, true},
			false,
		},
		{
			"Digmon no address",
			args{"WWG-FS001", "Test flow switch", "", false, true, true},
			&digmon{"WWG-FS001", "Test flow switch", "M0.0", false, true, true},
			false,
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
				tag:         "CTP-LSL001",
				description: "Test digmon 1",
				address:     "I3.0",
				invert:      false,
				alarm:       true,
				invertAlarm: false,
			},
			"CTP-LSL001",
		},
		{
			"Digmon 2",
			&digmon{
				tag:         "CTP-LSL002",
				description: "Test digmon 2",
				address:     "",
				invert:      true,
				alarm:       true,
				invertAlarm: false,
			},
			"CTP-LSL002",
		},
		{
			"Digmon 3",
			&digmon{
				tag:         "CTP-LSL003",
				description: "Test digmon 3",
				address:     "I3.1",
				invert:      false,
				alarm:       false,
				invertAlarm: false,
			},
			"CTP-LSL003",
		},
		{
			"Digmon 4",
			&digmon{
				tag:         "CTP-LSH004",
				description: "Test digmon 4",
				address:     "",
				invert:      false,
				alarm:       true,
				invertAlarm: true,
			},
			"CTP-LSH004",
		},
		{
			"Digmon 5",
			&digmon{
				tag:         "CTP-LSL005",
				description: "Test digmon 5",
				address:     "I3.2",
				invert:      true,
				alarm:       false,
				invertAlarm: false,
			},
			"CTP-LSL005",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Tag(); got != tt.want {
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
				tag:         "CTP-LSL001",
				description: "Test digmon 1",
				address:     "I3.0",
				invert:      false,
				alarm:       true,
				invertAlarm: false,
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
				tag:         "CTP-LSL002",
				description: "Test digmon 2",
				address:     "",
				invert:      true,
				alarm:       true,
				invertAlarm: false,
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
				tag:         "CTP-LSL003",
				description: "Test digmon 3",
				address:     "I3.1",
				invert:      false,
				alarm:       false,
				invertAlarm: false,
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
				tag:         "CTP-LSH004",
				description: "Test digmon 4",
				address:     "",
				invert:      false,
				alarm:       true,
				invertAlarm: true,
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
				tag:         "CTP-LSL005",
				description: "Test digmon 5",
				address:     "I3.2",
				invert:      true,
				alarm:       false,
				invertAlarm: false,
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
				tag:         "CTP-LSL001",
				description: "Test digmon 1",
				address:     "I3.0",
				invert:      false,
				alarm:       true,
				invertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL001", "Bool", "I3.0", "Test digmon 1"},
			},
		},
		{
			"Digmon 2",
			&digmon{
				tag:         "CTP-LSL002",
				description: "Test digmon 2",
				address:     "I3.1",
				invert:      true,
				alarm:       true,
				invertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL002", "Bool", "I3.1", "Test digmon 2"},
			},
		},
		{
			"Digmon 3",
			&digmon{
				tag:         "CTP-LSL003",
				description: "Test digmon 3",
				address:     "I3.1",
				invert:      false,
				alarm:       false,
				invertAlarm: false,
			},
			[]*PlcTag{
				{"CTP-LSL003", "Bool", "I3.1", "Test digmon 3"},
			},
		},
		{
			"Digmon 4",
			&digmon{
				tag:         "CTP-LSH004",
				description: "Test digmon 4",
				address:     "I3.7",
				invert:      false,
				alarm:       true,
				invertAlarm: true,
			},
			[]*PlcTag{
				{"CTP-LSH004", "Bool", "I3.7", "Test digmon 4"},
			},
		},
		{
			"Digmon 5",
			&digmon{
				tag:         "CTP-LSL005",
				description: "Test digmon 5",
				address:     "I3.2",
				invert:      true,
				alarm:       false,
				invertAlarm: false,
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
