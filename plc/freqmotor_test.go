package plc

import (
	"reflect"
	"testing"
)

func TestFreqMotor_Stringer(t *testing.T) {
	tests := []struct {
		name string
		f    *freqMotor
		want string
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I4.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I4.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I4.2",
				DanfossDrive:     false,
			},
			"WWG-P001",
		},
		{
			"No feedback",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackAddress:  "",
				BreakerAddress:   "I4.3",
				SwitchAddress:    "I4.4",
				DanfossDrive:     false,
			},
			"WWG-P002",
		},
		{
			"No breaker",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackAddress:  "I4.5",
				BreakerAddress:   "",
				SwitchAddress:    "I4.6",
				DanfossDrive:     false,
			},
			"WWG-P003",
		},
		{
			"No switch",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackAddress:  "I4.1",
				BreakerAddress:   "I5.0",
				SwitchAddress:    "",
				DanfossDrive:     false,
			},
			"WWG-P004",
		},
		{
			"No feedback & breaker",
			&freqMotor{
				Tag:              "WWG-P005",
				Description:      "Frequency motor 5",
				ContactorAddress: "Q4.4",
				PqwAddress:       "QW28",
				FeedbackAddress:  "",
				BreakerAddress:   "",
				SwitchAddress:    "I5.1",
				DanfossDrive:     false,
			},
			"WWG-P005",
		},
		{
			"Danfoss",
			&freqMotor{
				Tag:              "WWG-P006",
				Description:      "Frequency motor 6",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackAddress:  "I4.0",
				BreakerAddress:   "I4.1",
				SwitchAddress:    "I4.2",
				DanfossDrive:     true,
			},
			"WWG-P006",
		},
		{
			"No feedback danfoss",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackAddress:  "",
				BreakerAddress:   "I4.3",
				SwitchAddress:    "I4.4",
				DanfossDrive:     true,
			},
			"WWG-P002",
		},
		{
			"No breaker danfoss",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackAddress:  "I4.5",
				BreakerAddress:   "",
				SwitchAddress:    "I4.6",
				DanfossDrive:     true,
			},
			"WWG-P003",
		},
		{
			"No switch danfoss",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackAddress:  "I4.1",
				BreakerAddress:   "I5.0",
				SwitchAddress:    "",
				DanfossDrive:     true,
			},
			"WWG-P004",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Stringer(); got != tt.want {
				t.Errorf("FreqMotor.Stringer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreqMotor_InputMap(t *testing.T) {
	tests := []struct {
		name string
		f    *freqMotor
		want map[string]string
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I4.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I4.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I4.2",
				DanfossDrive:     false,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No feedback",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackAddress:  "",
				BreakerAddress:   "I4.3",
				SwitchAddress:    "I4.4",
				DanfossDrive:     false,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No breaker",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackAddress:  "I4.5",
				BreakerAddress:   "",
				SwitchAddress:    "I4.6",
				DanfossDrive:     false,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No switch",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackAddress:  "I4.1",
				BreakerAddress:   "I5.0",
				SwitchAddress:    "",
				DanfossDrive:     false,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No feedback & breaker",
			&freqMotor{
				Tag:              "WWG-P005",
				Description:      "Frequency motor 5",
				ContactorAddress: "Q4.4",
				PqwAddress:       "QW28",
				FeedbackAddress:  "",
				BreakerAddress:   "",
				SwitchAddress:    "I5.1",
				DanfossDrive:     false,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"Danfoss",
			&freqMotor{
				Tag:              "WWG-P006",
				Description:      "Frequency motor 6",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackAddress:  "I4.0",
				BreakerAddress:   "I4.1",
				SwitchAddress:    "I4.2",
				DanfossDrive:     true,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No feedback danfoss",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackAddress:  "",
				BreakerAddress:   "I4.3",
				SwitchAddress:    "I4.4",
				DanfossDrive:     true,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No breaker danfoss",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackAddress:  "I4.5",
				BreakerAddress:   "",
				SwitchAddress:    "I4.6",
				DanfossDrive:     true,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
		{
			"No switch danfoss",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackAddress:  "I4.1",
				BreakerAddress:   "I5.0",
				SwitchAddress:    "",
				DanfossDrive:     true,
			},
			map[string]string{
				"Tag":              "WWG-P001",
				"Description":      "Frequency motor 1",
				"ContactorAddress": "Q4.0",
				"PQWAddress":       "QW20",
				"FeedbackTag":      "WWG-P001_FB",
				"FeedbackAddress":  "I4.0",
				"BreakerTag":       "WWG-P001_TH",
				"BreakerAddress":   "I4.1",
				"SwitchTag":        "WWG-P001_WS",
				"SwitchAddress":    "I4.2",
				"Danfoss":          "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreqMotor.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreqMotor_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		f    *freqMotor
		want []PlcTag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.PlcTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreqMotor.PlcTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
