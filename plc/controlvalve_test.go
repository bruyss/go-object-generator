package plc

import (
	"reflect"
	"testing"
)

func TestControlValve_String(t *testing.T) {
	tests := []struct {
		name string
		c    *controlValve
		want string
	}{
		{
			"Case 1",
			&controlValve{
				tag:             "WWG-CV001",
				description:     "Test control valve 1",
				address:         "QW16",
				feedbackAddress: "IW32",
				monitoringTime:  10,
			},
			"WWG-CV001",
		},
		{
			"Case 2",
			&controlValve{
				tag:             "WWG-CV002",
				description:     "Test control valve 2",
				address:         "QW18",
				feedbackAddress: "IW34",
				monitoringTime:  20,
			},
			"WWG-CV002",
		},
		{
			"Case no feedback",
			&controlValve{
				tag:             "WWG-CV003",
				description:     "Test control valve 3",
				address:         "QW20",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			"WWG-CV003",
		},
		{
			"Case no address",
			&controlValve{
				tag:             "WWG-CV004",
				description:     "Test control valve 4",
				address:         "",
				feedbackAddress: "IW36",
				monitoringTime:  10,
			},
			"WWG-CV004",
		},
		{
			"Case no address, no feedback",
			&controlValve{
				tag:             "WWG-CV005",
				description:     "Test control valve 5",
				address:         "",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			"WWG-CV005",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Tag(); got != tt.want {
				t.Errorf("ControlValve.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestControlValve_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		m    *controlValve
		want []*PlcTag
	}{
		{
			"Case 1",
			&controlValve{
				tag:             "WWG-CV001",
				description:     "Test control valve 1",
				address:         "QW16",
				feedbackAddress: "IW32",
				monitoringTime:  10,
			},
			[]*PlcTag{
				{name: "WWG-CV001", dtype: "Int", address: "QW16", comment: "Test control valve 1 output"},
				{name: "WWG-CV001_FB", dtype: "Int", address: "IW32", comment: "Test control valve 1 feedback"},
			},
		},
		{
			"Case 2",
			&controlValve{
				tag:             "WWG-CV002",
				description:     "Test control valve 2",
				address:         "QW18",
				feedbackAddress: "IW34",
				monitoringTime:  20,
			},
			[]*PlcTag{
				{name: "WWG-CV002", dtype: "Int", address: "QW18", comment: "Test control valve 2 output"},
				{name: "WWG-CV002_FB", dtype: "Int", address: "IW34", comment: "Test control valve 2 feedback"},
			},
		},
		{
			"Case no feedback",
			&controlValve{
				tag:             "WWG-CV003",
				description:     "Test control valve 3",
				address:         "QW20",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			[]*PlcTag{
				{name: "WWG-CV003", dtype: "Int", address: "QW20", comment: "Test control valve 3 output"},
				{name: "WWG-CV003_FB", dtype: "Int", address: "MW2", comment: "Test control valve 3 feedback"},
			},
		},
		{
			"Case no address",
			&controlValve{
				tag:             "WWG-CV004",
				description:     "Test control valve 4",
				address:         "",
				feedbackAddress: "IW36",
				monitoringTime:  10,
			},
			[]*PlcTag{
				{name: "WWG-CV004", dtype: "Int", address: "MW0", comment: "Test control valve 4 output"},
				{name: "WWG-CV004_FB", dtype: "Int", address: "IW36", comment: "Test control valve 4 feedback"},
			},
		},
		{
			"Case no address, no feedback",
			&controlValve{
				tag:             "WWG-CV005",
				description:     "Test control valve 5",
				address:         "",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			[]*PlcTag{
				{name: "WWG-CV005", dtype: "Int", address: "MW0", comment: "Test control valve 5 output"},
				{name: "WWG-CV005_FB", dtype: "Int", address: "MW2", comment: "Test control valve 5 feedback"},
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

func TestControlValve_InputMap(t *testing.T) {
	tests := []struct {
		name string
		m    *controlValve
		want map[string]string
	}{
		{
			"Case 1",
			&controlValve{
				tag:             "WWG-CV001",
				description:     "Test control valve 1",
				address:         "QW16",
				feedbackAddress: "IW32",
				monitoringTime:  10,
			},
			map[string]string{
				"Tag":            "WWG-CV001",
				"Description":    "Test control valve 1",
				"IDB":            "IDB_WWG-CV001",
				"NoFeedback":     "false",
				"Feedback":       "WWG-CV001_FB",
				"MonitoringTime": "10",
				"Output":         `"WWG-CV001"`,
			},
		},
		{
			"Case 2",
			&controlValve{
				tag:             "WWG-CV002",
				description:     "Test control valve 2",
				address:         "QW18",
				feedbackAddress: "IW34",
				monitoringTime:  20,
			},
			map[string]string{
				"Tag":            "WWG-CV002",
				"Description":    "Test control valve 2",
				"IDB":            "IDB_WWG-CV002",
				"NoFeedback":     "false",
				"Feedback":       "WWG-CV002_FB",
				"MonitoringTime": "20",
				"Output":         `"WWG-CV002"`,
			},
		},
		{
			"Case no feedback",
			&controlValve{
				tag:             "WWG-CV003",
				description:     "Test control valve 3",
				address:         "QW20",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			map[string]string{
				"Tag":            "WWG-CV003",
				"Description":    "Test control valve 3",
				"IDB":            "IDB_WWG-CV003",
				"NoFeedback":     "true",
				"Feedback":       "WWG-CV003_FB",
				"MonitoringTime": "10",
				"Output":         `"WWG-CV003"`,
			},
		},
		{
			"Case no address",
			&controlValve{
				tag:             "WWG-CV004",
				description:     "Test control valve 4",
				address:         "",
				feedbackAddress: "IW36",
				monitoringTime:  10,
			},
			map[string]string{
				"Tag":            "WWG-CV004",
				"Description":    "Test control valve 4",
				"IDB":            "IDB_WWG-CV004",
				"NoFeedback":     "false",
				"Feedback":       "WWG-CV004_FB",
				"MonitoringTime": "10",
				"Output":         `"WWG-CV004"`,
			},
		},
		{
			"Case no address, no feedback",
			&controlValve{
				tag:             "WWG-CV005",
				description:     "Test control valve 5",
				address:         "",
				feedbackAddress: "",
				monitoringTime:  10,
			},
			map[string]string{
				"Tag":            "WWG-CV005",
				"Description":    "Test control valve 5",
				"IDB":            "IDB_WWG-CV005",
				"NoFeedback":     "true",
				"Feedback":       "WWG-CV005_FB",
				"MonitoringTime": "10",
				"Output":         `"WWG-CV005"`,
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
