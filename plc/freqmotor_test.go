package plc

import (
	"reflect"
	"testing"
)

func TestNewFreqMotor(t *testing.T) {
	type args struct {
		tag              string
		description      string
		contactorAddress string
		pqwAddress       string
		feedbackTag      string
		feedbackAddress  string
		breakerTag       string
		breakerAddress   string
		switchTag        string
		switchAddress    string
		danfossDrive     string
	}
	tests := []struct {
		name    string
		args    args
		want    *freqMotor
		wantErr bool
	}{
		{
			"Frequency motor",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no contactor address",
			args{"WWG-P001", "Frequency motor 1", "", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "M0.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no PQW address",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "MW0",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no feedback address",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "M0.1",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no breaker address",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "M0.2",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no switch address",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "M0.3",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no feedback",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "", "", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no breaker",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "", "", "WWG-P001_WS", "I1.2", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			false,
		},
		{
			"Frequency motor no switch",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "", "", "false"},
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "",
				SwitchAddress:    "",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			false,
		},
		{
			"Frequency motor bad danfoss",
			args{"WWG-P001", "Frequency motor 1", "Q1.0", "QW2", "WWG-P001_FB", "I1.0", "WWG-P001_TH", "I1.1", "WWG-P001_WS", "I1.2", "allo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFreqMotor(tt.args.tag, tt.args.description, tt.args.contactorAddress, tt.args.pqwAddress, tt.args.feedbackTag, tt.args.feedbackAddress, tt.args.breakerTag, tt.args.breakerAddress, tt.args.switchTag, tt.args.switchAddress, tt.args.danfossDrive)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFreqMotor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFreqMotor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqMotor_InputMap(t *testing.T) {
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
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			map[string]string{
				"Tag":          "WWG-P001",
				"Description":  "Frequency motor 1",
				"IDB":          "IDB_WWG-P001",
				"ContactorTag": `"WWG-P001"`,
				"PQW":          `"WWG-P001_PQW"`,
				"FeedbackTag":  `"WWG-P001_FB"`,
				"BreakerTag":   `"WWG-P001_TH"`,
				"SwitchTag":    `"WWG-P001_WS"`,
				"AlarmTag":     `"WWG-P001_AL"`,
				"Danfoss":      "false",
			},
		},
		{
			"No feedback",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-P002_TH",
				BreakerAddress:   "I4.3",
				SwitchTag:        "WWG-P002_WS",
				SwitchAddress:    "I4.4",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			map[string]string{
				"Tag":          "WWG-P002",
				"Description":  "Frequency motor 2",
				"IDB":          "IDB_WWG-P002",
				"ContactorTag": `"WWG-P002"`,
				"PQW":          `"WWG-P002_PQW"`,
				"FeedbackTag":  `"IDB_WWG-P002".Q_On`,
				"BreakerTag":   `"WWG-P002_TH"`,
				"SwitchTag":    `"WWG-P002_WS"`,
				"AlarmTag":     `"WWG-P002_AL"`,
				"Danfoss":      "false",
			},
		},
		{
			"No breaker",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackTag:      "WWG-P003_FB",
				FeedbackAddress:  "I4.5",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P003_WS",
				SwitchAddress:    "I4.6",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			map[string]string{
				"Tag":          "WWG-P003",
				"Description":  "Frequency motor 3",
				"IDB":          "IDB_WWG-P003",
				"ContactorTag": `"WWG-P003"`,
				"PQW":          `"WWG-P003_PQW"`,
				"FeedbackTag":  `"WWG-P003_FB"`,
				"BreakerTag":   "FALSE",
				"SwitchTag":    `"WWG-P003_WS"`,
				"AlarmTag":     `"WWG-P003_AL"`,
				"Danfoss":      "false",
			},
		},
		{
			"No switch",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackTag:      "WWG-P004_FB",
				FeedbackAddress:  "I4.1",
				BreakerTag:       "WWG-P004_TH",
				BreakerAddress:   "I5.0",
				SwitchTag:        "",
				SwitchAddress:    "",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			map[string]string{
				"Tag":          "WWG-P004",
				"Description":  "Frequency motor 4",
				"IDB":          "IDB_WWG-P004",
				"ContactorTag": `"WWG-P004"`,
				"PQW":          `"WWG-P004_PQW"`,
				"FeedbackTag":  `"WWG-P004_FB"`,
				"BreakerTag":   `"WWG-P004_TH"`,
				"SwitchTag":    "TRUE",
				"AlarmTag":     `"WWG-P004_AL"`,
				"Danfoss":      "false",
			},
		},
		{
			"No feedback & breaker",
			&freqMotor{
				Tag:              "WWG-P005",
				Description:      "Frequency motor 5",
				ContactorAddress: "Q4.4",
				PqwAddress:       "QW28",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P005_WS",
				SwitchAddress:    "I5.1",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			map[string]string{
				"Tag":          "WWG-P005",
				"Description":  "Frequency motor 5",
				"IDB":          "IDB_WWG-P005",
				"ContactorTag": `"WWG-P005"`,
				"PQW":          `"WWG-P005_PQW"`,
				"FeedbackTag":  `"IDB_WWG-P005".Q_On`,
				"BreakerTag":   "FALSE",
				"SwitchTag":    `"WWG-P005_WS"`,
				"AlarmTag":     `"WWG-P005_AL"`,
				"Danfoss":      "false",
			},
		},
		{
			"Danfoss",
			&freqMotor{
				Tag:              "WWG-P006",
				Description:      "Frequency motor 6",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackTag:      "WWG-P006_FB",
				FeedbackAddress:  "I4.0",
				BreakerTag:       "WWG-P006_TH",
				BreakerAddress:   "I4.1",
				SwitchTag:        "WWG-P006_WS",
				SwitchAddress:    "I4.2",
				DanfossDrive:     true,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			map[string]string{
				"Tag":          "WWG-P006",
				"Description":  "Frequency motor 6",
				"IDB":          "IDB_WWG-P006",
				"ContactorTag": `"WWG-P006"`,
				"PQW":          `"WWG-P006_PQW"`,
				"FeedbackTag":  `"WWG-P006_FB"`,
				"BreakerTag":   `"WWG-P006_TH"`,
				"SwitchTag":    `"WWG-P006_WS"`,
				"AlarmTag":     `"WWG-P006_AL"`,
				"Danfoss":      "true",
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

func Test_freqMotor_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		f    *freqMotor
		want []*PlcTag
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
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P001", "Bool", "Q4.0", "Frequency motor 1"},
				{"WWG-P001_PQW", "Int", "QW20", "Frequency motor 1 output"},
				{"WWG-P001_FB", "Bool", "I4.0", "Frequency motor 1 feedback"},
				{"WWG-P001_TH", "Bool", "I4.1", "Frequency motor 1 breaker"},
				{"WWG-P001_WS", "Bool", "I4.2", "Frequency motor 1 protection switch"},
			},
		},
		{
			"No feedback",
			&freqMotor{
				Tag:              "WWG-P002",
				Description:      "Frequency motor 2",
				ContactorAddress: "Q4.1",
				PqwAddress:       "QW22",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-P002_TH",
				BreakerAddress:   "I4.3",
				SwitchTag:        "WWG-P002_WS",
				SwitchAddress:    "I4.4",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P002", "Bool", "Q4.1", "Frequency motor 2"},
				{"WWG-P002_PQW", "Int", "QW22", "Frequency motor 2 output"},
				{"WWG-P002_TH", "Bool", "I4.3", "Frequency motor 2 breaker"},
				{"WWG-P002_WS", "Bool", "I4.4", "Frequency motor 2 protection switch"},
			},
		},
		{
			"No breaker",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackTag:      "WWG-P003_FB",
				FeedbackAddress:  "I4.5",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P003_WS",
				SwitchAddress:    "I4.6",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P003", "Bool", "Q4.2", "Frequency motor 3"},
				{"WWG-P003_PQW", "Int", "QW24", "Frequency motor 3 output"},
				{"WWG-P003_FB", "Bool", "I4.5", "Frequency motor 3 feedback"},
				{"WWG-P003_WS", "Bool", "I4.6", "Frequency motor 3 protection switch"},
			},
		},
		{
			"No switch",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackTag:      "WWG-P004_FB",
				FeedbackAddress:  "I4.1",
				BreakerTag:       "WWG-P004_TH",
				BreakerAddress:   "I5.0",
				SwitchTag:        "",
				SwitchAddress:    "",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			[]*PlcTag{
				{"WWG-P004", "Bool", "Q4.3", "Frequency motor 4"},
				{"WWG-P004_PQW", "Int", "QW26", "Frequency motor 4 output"},
				{"WWG-P004_FB", "Bool", "I4.1", "Frequency motor 4 feedback"},
				{"WWG-P004_TH", "Bool", "I5.0", "Frequency motor 4 breaker"},
			},
		},
		{
			"No feedback & breaker",
			&freqMotor{
				Tag:              "WWG-P005",
				Description:      "Frequency motor 5",
				ContactorAddress: "Q4.4",
				PqwAddress:       "QW28",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P005_WS",
				SwitchAddress:    "I5.1",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P005", "Bool", "Q4.4", "Frequency motor 5"},
				{"WWG-P005_PQW", "Int", "QW28", "Frequency motor 5 output"},
				{"WWG-P005_WS", "Bool", "I5.1", "Frequency motor 5 protection switch"},
			},
		},
		{
			"Danfoss",
			&freqMotor{
				Tag:              "WWG-P006",
				Description:      "Frequency motor 6",
				ContactorAddress: "Q4.0",
				PqwAddress:       "QW20",
				FeedbackTag:      "WWG-P006_FB",
				FeedbackAddress:  "I4.0",
				BreakerTag:       "WWG-P006_TH",
				BreakerAddress:   "I4.1",
				SwitchTag:        "WWG-P006_WS",
				SwitchAddress:    "I4.2",
				DanfossDrive:     true,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P006_TH", "Bool", "I4.1", "Frequency motor 6 breaker"},
				{"WWG-P006_WS", "Bool", "I4.2", "Frequency motor 6 protection switch"},
			},
		},
		{
			"No breaker danfoss",
			&freqMotor{
				Tag:              "WWG-P003",
				Description:      "Frequency motor 3",
				ContactorAddress: "Q4.2",
				PqwAddress:       "QW24",
				FeedbackTag:      "WWG-P003_FB",
				FeedbackAddress:  "I4.5",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-P003_WS",
				SwitchAddress:    "I4.6",
				DanfossDrive:     true,
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			[]*PlcTag{
				{"WWG-P003_WS", "Bool", "I4.6", "Frequency motor 3 protection switch"},
			},
		},
		{
			"No switch danfoss",
			&freqMotor{
				Tag:              "WWG-P004",
				Description:      "Frequency motor 4",
				ContactorAddress: "Q4.3",
				PqwAddress:       "QW26",
				FeedbackTag:      "",
				FeedbackAddress:  "I4.1",
				BreakerTag:       "WWG-P004_TH",
				BreakerAddress:   "I5.0",
				SwitchTag:        "",
				SwitchAddress:    "",
				DanfossDrive:     true,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			[]*PlcTag{
				{"WWG-P004_TH", "Bool", "I5.0", "Frequency motor 4 breaker"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.PlcTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreqMotor.PlcTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqMotor_contactorPlcTag(t *testing.T) {
	tests := []struct {
		name  string
		f     *freqMotor
		wantT *PlcTag
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			&PlcTag{Name: "WWG-P001", Dtype: "Bool", Address: "Q1.0", Comment: "Frequency motor 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.f.contactorPlcTag(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("freqMotor.contactorPlcTag() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_freqMotor_pqwPlcTag(t *testing.T) {
	tests := []struct {
		name  string
		f     *freqMotor
		wantT *PlcTag
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			&PlcTag{Name: "WWG-P001_PQW", Dtype: "Int", Address: "QW2", Comment: "Frequency motor 1 output"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.f.pqwPlcTag(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("freqMotor.pqwPlcTag() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_freqMotor_feedbackPlcTag(t *testing.T) {
	tests := []struct {
		name  string
		f     *freqMotor
		wantT *PlcTag
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			&PlcTag{Name: "WWG-P001_FB", Dtype: "Bool", Address: "I1.0", Comment: "Frequency motor 1 feedback"},
		},
		{
			"Frequency motor no feedback",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			nil,
		},
		{
			"Frequency motor danfoss",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     true,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.f.feedbackPlcTag(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("freqMotor.feedbackPlcTag() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_freqMotor_breakerPlcTag(t *testing.T) {
	tests := []struct {
		name  string
		f     *freqMotor
		wantT *PlcTag
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			&PlcTag{Name: "WWG-P001_TH", Dtype: "Bool", Address: "I1.1", Comment: "Frequency motor 1 breaker"},
		},
		{
			"Frequency motor no feedback",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.f.breakerPlcTag(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("freqMotor.breakerPlcTag() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_freqMotor_switchPlcTag(t *testing.T) {
	tests := []struct {
		name  string
		f     *freqMotor
		wantT *PlcTag
	}{
		{
			"Frequency motor",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			&PlcTag{Name: "WWG-P001_WS", Dtype: "Bool", Address: "I1.2", Comment: "Frequency motor 1 protection switch"},
		},
		{
			"Frequency motor no feedback",
			&freqMotor{
				Tag:              "WWG-P001",
				Description:      "Frequency motor 1",
				ContactorAddress: "Q1.0",
				PqwAddress:       "QW2",
				FeedbackTag:      "WWG-P001_FB",
				FeedbackAddress:  "I1.0",
				BreakerTag:       "WWG-P001_TH",
				BreakerAddress:   "I1.1",
				SwitchTag:        "WWG-P001_WS",
				SwitchAddress:    "I1.2",
				DanfossDrive:     false,
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.f.switchPlcTag(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("freqMotor.switchPlcTag() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
