package plc

import (
	"reflect"
	"testing"
)

func TestNewDigout(t *testing.T) {
	type args struct {
		tag             string
		description     string
		outputAddress   string
		feedbackTag     string
		feedbackAddress string
		breakerTag      string
		breakerAddress  string
		monitoringTime  string
		data            map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *digout
		wantErr bool
	}{
		{
			name: "Digital out",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Digital out extra data",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data: map[string]string{
					"Custom 1": "data 1",
					"Custom 2": "data 2",
				},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data: map[string]string{
					"Custom 1": "data 1",
					"Custom 2": "data 2",
				},
			},
			wantErr: false,
		},
		{
			name: "Digital out no output address",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data: map[string]string{
					"Custom 1": "data 1",
					"Custom 2": "data 2",
				},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "M0.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data: map[string]string{
					"Custom 1": "data 1",
					"Custom 2": "data 2",
				},
			},
			wantErr: false,
		},
		{
			name: "Digital out no feedback tag",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "",
				FeedbackAddress: "",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     false,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Digital out no feedback address",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "M0.1",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Digital out no breaker tag",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "",
				breakerAddress:  "I2.1",
				monitoringTime:  "15",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "",
				BreakerAddress:  "",
				hasFeedback:     true,
				hasBreaker:      false,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Digital out no breaker address",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "",
				monitoringTime:  "15",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "M0.2",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Digital out monitoring time",
			args: args{
				tag:             "WWG-D001",
				description:     "Test digital output 1",
				outputAddress:   "Q2.0",
				feedbackTag:     "WWG-D001_FB",
				feedbackAddress: "I2.0",
				breakerTag:      "WWG-D001_TH",
				breakerAddress:  "I2.1",
				monitoringTime:  "abc",
				data:            map[string]string{},
			},
			want: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  10,
				Data:            map[string]string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDigout(tt.args.tag, tt.args.description, tt.args.outputAddress, tt.args.feedbackTag, tt.args.feedbackAddress, tt.args.breakerTag, tt.args.breakerAddress, tt.args.monitoringTime, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDigout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDigout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digout_InputMap(t *testing.T) {
	tests := []struct {
		name string
		d    *digout
		want map[string]string
	}{
		{
			name: "Digital out",
			d: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			want: map[string]string{
				"Tag":            "WWG-D001",
				"Description":    "Test digital output 1",
				"IDB":            "IDB_WWG-D001",
				"OutputTag":      `"WWG-D001"`,
				"FeedbackTag":    `"WWG-D001_FB"`,
				"BreakerTag":     `"WWG-D001_TH"`,
				"MonitoringTime": "15",
			},
		},
		{
			name: "Digital out extra data",
			d: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "WWG-D001_FB",
				FeedbackAddress: "I2.0",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     true,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data: map[string]string{
					"Custom 1": "data 1",
					"Custom 2": "data 2",
				},
			},
			want: map[string]string{
				"Tag":            "WWG-D001",
				"Description":    "Test digital output 1",
				"IDB":            "IDB_WWG-D001",
				"OutputTag":      `"WWG-D001"`,
				"FeedbackTag":    `"WWG-D001_FB"`,
				"BreakerTag":     `"WWG-D001_TH"`,
				"MonitoringTime": "15",
				"Custom 1":       "data 1",
				"Custom 2":       "data 2",
			},
		},
		{
			name: "Digital out no feedback tag",
			d: &digout{
				Tag:             "WWG-D001",
				Description:     "Test digital output 1",
				OutputAddress:   "Q2.0",
				FeedbackTag:     "",
				FeedbackAddress: "",
				BreakerTag:      "WWG-D001_TH",
				BreakerAddress:  "I2.1",
				hasFeedback:     false,
				hasBreaker:      true,
				MonitoringTime:  15,
				Data:            map[string]string{},
			},
			want: map[string]string{
				"Tag":            "WWG-D001",
				"Description":    "Test digital output 1",
				"IDB":            "IDB_WWG-D001",
				"OutputTag":      `"WWG-D001"`,
				"FeedbackTag":    `"IDB_WWG-D001".Q_On`,
				"BreakerTag":     `"WWG-D001_TH"`,
				"MonitoringTime": "15",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digout.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digout_PlcTags(t *testing.T) {
	tests := []struct {
		name  string
		d     *digout
		wantT []*PlcTag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.d.PlcTags(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("digout.PlcTags() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_digout_outputTag(t *testing.T) {
	tests := []struct {
		name string
		d    *digout
		want *PlcTag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.outputTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digout.outputTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digout_feedbackTag(t *testing.T) {
	tests := []struct {
		name string
		d    *digout
		want *PlcTag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.feedbackTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digout.feedbackTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digout_breakerTag(t *testing.T) {
	tests := []struct {
		name string
		d    *digout
		want *PlcTag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.breakerTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digout.breakerTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
