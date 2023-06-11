package plc

import (
	"reflect"
	"testing"

	"github.com/bruyss/go-object-generator/logger"
)

func init() {
	logger.InitializeDevLogger()
}

/*  */
func TestNewControlValve(t *testing.T) {
	type args struct {
		tag             string
		description     string
		address         string
		feedbackTag     string
		feedbackAddress string
		monitoringTime  string
		data            map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *controlValve
		wantErr bool
	}{
		{
			"Control valve",
			args{"WWG-CV001", "Control valve 1", "QW10", "WWG-CV001_FB", "IW32", "14", map[string]string{}},
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Control valve 1",
				Address:         "QW10",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  14,
				hasFeedback:     true,
				Data:            map[string]string{},
			},
			false,
		},
		{
			"Control valve extra data",
			args{"WWG-CV001", "Control valve 1", "QW10", "WWG-CV001_FB", "IW32", "14", map[string]string{"Custom 1": "allo"}},
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Control valve 1",
				Address:         "QW10",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  14,
				hasFeedback:     true,
				Data:            map[string]string{"Custom 1": "allo"},
			},
			false,
		},
		{
			"Control valve no address",
			args{"WWG-CV001", "Control valve 1", "", "WWG-CV001_FB", "IW32", "14", map[string]string{}},
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Control valve 1",
				Address:         "MW0",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  14,
				hasFeedback:     true,
				Data:            map[string]string{},
			},
			false,
		},
		{
			"Control valve no feedback",
			args{"WWG-CV001", "Control valve 1", "QW10", "", "", "14", map[string]string{}},
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Control valve 1",
				Address:         "QW10",
				FeedbackTag:     "",
				FeedbackAddress: "",
				MonitoringTime:  14,
				hasFeedback:     false,
				Data:            map[string]string{},
			},
			false,
		},
		{
			"Control valve no feedback address",
			args{"WWG-CV001", "Control valve 1", "QW10", "WWG-CV001_FB", "", "14", map[string]string{}},
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Control valve 1",
				Address:         "QW10",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "MW2",
				MonitoringTime:  14,
				hasFeedback:     true,
				Data:            map[string]string{},
			},
			false,
		},
		{
			"Control valve bad monitoring time",
			args{"WWG-CV001", "Control valve 1", "QW10", "WWG-CV001_FB", "IW32", "allo", map[string]string{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewControlValve(tt.args.tag, tt.args.description, tt.args.address, tt.args.feedbackTag, tt.args.feedbackAddress, tt.args.monitoringTime, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewControlValve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewControlValve() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_controlValve_PlcTags(t *testing.T) {
	tests := []struct {
		name string
		m    *controlValve
		want []*PlcTag
	}{
		{
			"Control valve",
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackTag:     "",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
				hasFeedback:     true,
			},
			[]*PlcTag{
				{Name: "WWG-CV001", Dtype: "Int", Address: "QW16", Comment: "Test control valve 1 output"},
				{Name: "WWG-CV001_FB", Dtype: "Int", Address: "IW32", Comment: "Test control valve 1 feedback"},
			},
		},
		{
			"Control valve 2",
			&controlValve{
				Tag:             "WWG-CV002",
				Description:     "Test control valve 2",
				Address:         "QW18",
				FeedbackTag:     "",
				FeedbackAddress: "IW34",
				MonitoringTime:  20,
				hasFeedback:     true,
			},
			[]*PlcTag{
				{Name: "WWG-CV002", Dtype: "Int", Address: "QW18", Comment: "Test control valve 2 output"},
				{Name: "WWG-CV002_FB", Dtype: "Int", Address: "IW34", Comment: "Test control valve 2 feedback"},
			},
		},
		{
			"Control valve no feedback",
			&controlValve{
				Tag:             "WWG-CV003",
				Description:     "Test control valve 3",
				Address:         "QW20",
				FeedbackTag:     "",
				FeedbackAddress: "",
				MonitoringTime:  10,
				hasFeedback:     false,
			},
			[]*PlcTag{
				{Name: "WWG-CV003", Dtype: "Int", Address: "QW20", Comment: "Test control valve 3 output"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.PlcTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controlValve.PlcTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controlValve_InputMap(t *testing.T) {
	tests := []struct {
		name string
		m    *controlValve
		want map[string]string
	}{
		{
			"Case 1",
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
				hasFeedback:     true,
			},
			map[string]string{
				"Tag":            "WWG-CV001",
				"Description":    "Test control valve 1",
				"IDB":            "IDB_WWG-CV001",
				"NoFeedback":     "FALSE",
				"Feedback":       `"WWG-CV001_FB"`,
				"MonitoringTime": "10",
				"Output":         `"WWG-CV001"`,
			},
		},
		{
			"Case 1 extra data",
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
				hasFeedback:     true,
				Data:            map[string]string{"Custom 1": "allo", "Tag": "dont"},
			},
			map[string]string{
				"Tag":            "WWG-CV001",
				"Description":    "Test control valve 1",
				"IDB":            "IDB_WWG-CV001",
				"NoFeedback":     "FALSE",
				"Feedback":       `"WWG-CV001_FB"`,
				"MonitoringTime": "10",
				"Output":         `"WWG-CV001"`,
				"Custom 1":       "allo",
			},
		},
		{
			"Case 2",
			&controlValve{
				Tag:             "WWG-CV002",
				Description:     "Test control valve 2",
				Address:         "QW18",
				FeedbackTag:     "WWG-CV002_FB",
				FeedbackAddress: "IW34",
				MonitoringTime:  20,
				hasFeedback:     true,
			},
			map[string]string{
				"Tag":            "WWG-CV002",
				"Description":    "Test control valve 2",
				"IDB":            "IDB_WWG-CV002",
				"NoFeedback":     "FALSE",
				"Feedback":       `"WWG-CV002_FB"`,
				"MonitoringTime": "20",
				"Output":         `"WWG-CV002"`,
			},
		},
		{
			"Case no feedback",
			&controlValve{
				Tag:             "WWG-CV003",
				Description:     "Test control valve 3",
				Address:         "QW20",
				FeedbackTag:     "",
				FeedbackAddress: "",
				MonitoringTime:  10,
				hasFeedback:     false,
			},
			map[string]string{
				"Tag":            "WWG-CV003",
				"Description":    "Test control valve 3",
				"IDB":            "IDB_WWG-CV003",
				"NoFeedback":     "TRUE",
				"Feedback":       `"IDB_WWG-CV003".Q_On`,
				"MonitoringTime": "10",
				"Output":         `"WWG-CV003"`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controlValve.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controlValve_outputPlcTag(t *testing.T) {
	tests := []struct {
		name string
		c    *controlValve
		want *PlcTag
	}{
		{
			"Control valve",
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
				hasFeedback:     true,
			},
			&PlcTag{"WWG-CV001", "Int", "QW16", "Test control valve 1 output"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.outputPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controlValve.outputPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controlValve_feedbackPlcTag(t *testing.T) {
	tests := []struct {
		name string
		c    *controlValve
		want *PlcTag
	}{
		{
			"Control valve",
			&controlValve{
				Tag:             "WWG-CV001",
				Description:     "Test control valve 1",
				Address:         "QW16",
				FeedbackTag:     "WWG-CV001_FB",
				FeedbackAddress: "IW32",
				MonitoringTime:  10,
				hasFeedback:     true,
			},
			&PlcTag{"WWG-CV001_FB", "Int", "IW32", "Test control valve 1 feedback"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.feedbackPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("controlValve.feedbackPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
