package plc

import (
	"reflect"
	"testing"
)

func TestControlValve_Stringer(t *testing.T) {
	tests := []struct {
		name string
		c    *ControlValve
		want string
	}{
		{
			"Case 1",
			&ControlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
			},
			"WWG-CV001",
		},
		{
			"Case 2",
			&ControlValve{
				Tag:             "WWG-CV002",
				Description:     "Test control valve 2",
				Address:         "QW18",
				FeedbackAddress: "IW34",
				MonitoringTime:  20,
			},
			"WWG-CV002",
		},
		{
			"Case no feedback",
			&ControlValve{
				Tag:             "WWG-CV003",
				Description:     "Test control valve 3",
				Address:         "QW20",
				FeedbackAddress: "",
				MonitoringTime:  10,
			},
			"WWG-CV003",
		},
		{
			"Case no address",
			&ControlValve{
				Tag:             "WWG-CV004",
				Description:     "Test control valve 4",
				Address:         "",
				FeedbackAddress: "IW36",
				MonitoringTime:  10,
			},
			"WWG-CV004",
		},
		{
			"Case no address, no feedback",
			&ControlValve{
				Tag:             "WWG-CV005",
				Description:     "Test control valve 5",
				Address:         "",
				FeedbackAddress: "",
				MonitoringTime:  10,
			},
			"WWG-CV005",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Stringer(); got != tt.want {
				t.Errorf("ControlValve.Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestControlValve_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		m    *ControlValve
		want []PlcTag
	}{
		{
			"Case 1",
			&ControlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
			},
			[]PlcTag{
				{name: "WWG-CV001", dtype: "Int", address: "QW16", comment: "Test control valve 1 output"},
				{name: "WWG-CV001_FB", dtype: "Int", address: "IW32", comment: "Test control valve 1 feedback"},
			},
		},
		{
			"Case 2",
			&ControlValve{
				Tag:             "WWG-CV002",
				Description:     "Test control valve 2",
				Address:         "QW18",
				FeedbackAddress: "IW34",
				MonitoringTime:  20,
			},
			[]PlcTag{
				{name: "WWG-CV002", dtype: "Int", address: "QW18", comment: "Test control valve 2 output"},
				{name: "WWG-CV002_FB", dtype: "Int", address: "IW34", comment: "Test control valve 2 feedback"},
			},
		},
		{
			"Case no feedback",
			&ControlValve{
				Tag:             "WWG-CV003",
				Description:     "Test control valve 3",
				Address:         "QW20",
				FeedbackAddress: "",
				MonitoringTime:  10,
			},
			[]PlcTag{
				{name: "WWG-CV003", dtype: "Int", address: "QW20", comment: "Test control valve 3 output"},
				{name: "WWG-CV003_FB", dtype: "Int", address: "MW2", comment: "Test control valve 3 feedback"},
			},
		},
		{
			"Case no address",
			&ControlValve{
				Tag:             "WWG-CV004",
				Description:     "Test control valve 4",
				Address:         "",
				FeedbackAddress: "IW36",
				MonitoringTime:  10,
			},
			[]PlcTag{
				{name: "WWG-CV004", dtype: "Int", address: "MW0", comment: "Test control valve 4 output"},
				{name: "WWG-CV004_FB", dtype: "Int", address: "IW36", comment: "Test control valve 4 feedback"},
			},
		},
		{
			"Case no address, no feedback",
			&ControlValve{
				Tag:             "WWG-CV005",
				Description:     "Test control valve 5",
				Address:         "",
				FeedbackAddress: "",
				MonitoringTime:  10,
			},
			[]PlcTag{
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
		m    *ControlValve
		want map[string]string
	}{
		{
			"Case 1",
			&ControlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
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
			&ControlValve{
				Tag:             "WWG-CV002",
				Description:     "Test control valve 2",
				Address:         "QW18",
				FeedbackAddress: "IW34",
				MonitoringTime:  20,
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
			&ControlValve{
				Tag:             "WWG-CV003",
				Description:     "Test control valve 3",
				Address:         "QW20",
				FeedbackAddress: "",
				MonitoringTime:  10,
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
			&ControlValve{
				Tag:             "WWG-CV004",
				Description:     "Test control valve 4",
				Address:         "",
				FeedbackAddress: "IW36",
				MonitoringTime:  10,
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
			&ControlValve{
				Tag:             "WWG-CV005",
				Description:     "Test control valve 5",
				Address:         "",
				FeedbackAddress: "",
				MonitoringTime:  10,
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
