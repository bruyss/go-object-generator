package plc

import (
	"reflect"
	"testing"
)

func TestNewMotor(t *testing.T) {
	type args struct {
		tag              string
		description      string
		contactorAddress string
		feedbackTag      string
		feedbackAddress  string
		breakerTag       string
		breakerAddress   string
		switchTag        string
		switchAddress    string
	}
	tests := []struct {
		name    string
		args    args
		want    *motor
		wantErr bool
	}{
		{
			name: "Motor",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no contactor",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "M0.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no feedback address",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "M0.1",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no breaker address",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "M0.2",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no switch address",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "M0.3",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no feedback",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "",
				feedbackAddress:  "",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no breaker",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "",
				breakerAddress:   "",
				switchTag:        "WWG-M001_WS",
				switchAddress:    "I10.2",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			wantErr: false,
		},
		{
			name: "Motor no switch",
			args: args{
				tag:              "WWG-M001",
				description:      "Test motor 1",
				contactorAddress: "Q10.0",
				feedbackTag:      "WWG-M001_FB",
				feedbackAddress:  "I10.0",
				breakerTag:       "WWG-M001_TH",
				breakerAddress:   "I10.1",
				switchTag:        "",
				switchAddress:    "",
			},
			want: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "",
				SwitchAddress:    "",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMotor(tt.args.tag, tt.args.description, tt.args.contactorAddress, tt.args.feedbackTag, tt.args.feedbackAddress, tt.args.breakerTag, tt.args.breakerAddress, tt.args.switchTag, tt.args.switchAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMotor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMotor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_String(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want string
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: `{"Tag":"WWG-M001","Description":"Test motor 1","ContactorAddress":"Q10.0","FeedbackTag":"WWG-M001_FB","FeedbackAddress":"I10.0","BreakerTag":"WWG-M001_TH","BreakerAddress":"I10.1","SwitchTag":"WWG-M001_WS","SwitchAddress":"I10.2"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("motor.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_InputMap(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want map[string]string
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: map[string]string{
				"Tag":          "WWG-M001",
				"Description":  "Test motor 1",
				"IDB":          `"IDB_WWG-M001"`,
				"ContactorTag": `"WWG-M001"`,
				"FeedbackTag":  `"WWG-M001_FB"`,
				"BreakerTag":   `"WWG-M001_TH"`,
				"SwitchTag":    `"WWG-M001_WS"`,
			},
		},
		{
			name: "Motor no feedback",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: map[string]string{
				"Tag":          "WWG-M001",
				"Description":  "Test motor 1",
				"IDB":          `"IDB_WWG-M001"`,
				"ContactorTag": `"WWG-M001"`,
				"FeedbackTag":  `"IDB_WWG-M001".Q_On`,
				"BreakerTag":   `"WWG-M001_TH"`,
				"SwitchTag":    `"WWG-M001_WS"`,
			},
		},
		{
			name: "Motor no breaker",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			want: map[string]string{
				"Tag":          "WWG-M001",
				"Description":  "Test motor 1",
				"IDB":          `"IDB_WWG-M001"`,
				"ContactorTag": `"WWG-M001"`,
				"FeedbackTag":  `"WWG-M001_FB"`,
				"BreakerTag":   "true",
				"SwitchTag":    `"WWG-M001_WS"`,
			},
		},
		{
			name: "Motor no switch",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "",
				SwitchAddress:    "",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			want: map[string]string{
				"Tag":          "WWG-M001",
				"Description":  "Test motor 1",
				"IDB":          `"IDB_WWG-M001"`,
				"ContactorTag": `"WWG-M001"`,
				"FeedbackTag":  `"WWG-M001_FB"`,
				"BreakerTag":   `"WWG-M001_TH"`,
				"SwitchTag":    "false",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("motor.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_PlcTags(t *testing.T) {
	tests := []struct {
		name  string
		m     *motor
		wantT []*PlcTag
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantT: []*PlcTag{
				{name: "WWG-M001", dtype: "Bool", address: "Q10.0", comment: "Test motor 1"},
				{name: "WWG-M001_FB", dtype: "Bool", address: "I10.0", comment: "Test motor 1 feedback"},
				{name: "WWG-M001_TH", dtype: "Bool", address: "I10.1", comment: "Test motor 1 breaker"},
				{name: "WWG-M001_WS", dtype: "Bool", address: "I10.2", comment: "Test motor 1 protection switch"},
			},
		},
		{
			name: "Motor no feedback",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			wantT: []*PlcTag{
				{name: "WWG-M001", dtype: "Bool", address: "Q10.0", comment: "Test motor 1"},
				{name: "WWG-M001_TH", dtype: "Bool", address: "I10.1", comment: "Test motor 1 breaker"},
				{name: "WWG-M001_WS", dtype: "Bool", address: "I10.2", comment: "Test motor 1 protection switch"},
			},
		},
		{
			name: "Motor no breaker",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			wantT: []*PlcTag{
				{name: "WWG-M001", dtype: "Bool", address: "Q10.0", comment: "Test motor 1"},
				{name: "WWG-M001_FB", dtype: "Bool", address: "I10.0", comment: "Test motor 1 feedback"},
				{name: "WWG-M001_WS", dtype: "Bool", address: "I10.2", comment: "Test motor 1 protection switch"},
			},
		},
		{
			name: "Motor no switch",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "",
				SwitchAddress:    "",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			wantT: []*PlcTag{
				{name: "WWG-M001", dtype: "Bool", address: "Q10.0", comment: "Test motor 1"},
				{name: "WWG-M001_FB", dtype: "Bool", address: "I10.0", comment: "Test motor 1 feedback"},
				{name: "WWG-M001_TH", dtype: "Bool", address: "I10.1", comment: "Test motor 1 breaker"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.m.PlcTags(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("motor.PlcTags() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_motor_contactorPlcTag(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want *PlcTag
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: &PlcTag{
				name:    "WWG-M001",
				dtype:   "Bool",
				address: "Q10.0",
				comment: "Test motor 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.contactorPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("motor.contactorPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_feedbackPlcTag(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want *PlcTag
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: &PlcTag{
				name:    "WWG-M001_FB",
				dtype:   "Bool",
				address: "I10.0",
				comment: "Test motor 1 feedback",
			},
		},
		{
			name: "Motor no feedback",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "",
				FeedbackAddress:  "",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      false,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.feedbackPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("motor.feedbackPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_breakerPlcTag(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want *PlcTag
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: &PlcTag{
				name:    "WWG-M001_TH",
				dtype:   "Bool",
				address: "I10.1",
				comment: "Test motor 1 breaker",
			},
		},
		{
			name: "Motor no breaker",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "",
				BreakerAddress:   "",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       false,
				hasSwitch:        true,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.breakerPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("motor.breakerPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_motor_switchPlcTag(t *testing.T) {
	tests := []struct {
		name string
		m    *motor
		want *PlcTag
	}{
		{
			name: "Motor",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "WWG-M001_WS",
				SwitchAddress:    "I10.2",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        true,
			},
			want: &PlcTag{
				name:    "WWG-M001_WS",
				dtype:   "Bool",
				address: "I10.2",
				comment: "Test motor 1 protection switch",
			},
		},
		{
			name: "Motor no switch",
			m: &motor{
				Tag:              "WWG-M001",
				Description:      "Test motor 1",
				ContactorAddress: "Q10.0",
				FeedbackTag:      "WWG-M001_FB",
				FeedbackAddress:  "I10.0",
				BreakerTag:       "WWG-M001_TH",
				BreakerAddress:   "I10.1",
				SwitchTag:        "",
				SwitchAddress:    "",
				hasFeedback:      true,
				hasBreaker:       true,
				hasSwitch:        false,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.switchPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("motor.switchPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
