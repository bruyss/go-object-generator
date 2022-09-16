package plc

import (
	"reflect"
	"testing"
)

func TestNewValve(t *testing.T) {
	type args struct {
		tag          string
		description  string
		actAddress   string
		fboTag       string
		fbcTag       string
		fboAddress   string
		fbcAddress   string
		monTimeOpen  string
		monTimeClose string
	}
	tests := []struct {
		name    string
		args    args
		want    *valve
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewValve(tt.args.tag, tt.args.description, tt.args.actAddress, tt.args.fboTag, tt.args.fbcTag, tt.args.fboAddress, tt.args.fbcAddress, tt.args.monTimeOpen, tt.args.monTimeClose)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewValve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValve() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_valve_String(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want string
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.0",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I0.1",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			`{"Tag":"WWG-XV001","Description":"Test valve 1","ActAddress":"Q0.0","FboTag":"WWG-XV001_FBO","FbcTag":"WWG-XV001_FBC","FboAddress":"I0.1","FbcAddress":"I0.2","MonTimeOpen":10,"MonTimeClose":15}`,
		},
		{
			"Valve no feedback open",
			&valve{
				Tag:          "WWG-XV002",
				Description:  "Test valve 2",
				ActAddress:   "Q0.1",
				FboTag:       "",
				FbcTag:       "WWG-XV002_FBC",
				FboAddress:   "",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			`{"Tag":"WWG-XV002","Description":"Test valve 2","ActAddress":"Q0.1","FboTag":"","FbcTag":"WWG-XV002_FBC","FboAddress":"","FbcAddress":"I0.2","MonTimeOpen":10,"MonTimeClose":10}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("valve.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valve_InputMap(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want map[string]string
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.0",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I0.1",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV001",
				"Description":  "Test valve 1",
				"IDB":          "IDB_WWG-XV001",
				"Output":       `"WWG-XV001"`,
				"FBO":          `"WWG-XV001_FBO"`,
				"FBC":          `"WWG-XV001_FBC"`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback open",
			&valve{
				Tag:          "WWG-XV002",
				Description:  "Test valve 2",
				ActAddress:   "Q0.1",
				FboTag:       "",
				FbcTag:       "WWG-XV002_FBC",
				FboAddress:   "",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV002",
				"Description":  "Test valve 2",
				"IDB":          "IDB_WWG-XV002",
				"Output":       `"WWG-XV002"`,
				"FBO":          `"IDB_WWG-XV002".Q_On`,
				"FBC":          `"WWG-XV002_FBC"`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "10",
			},
		},
		{
			"Valve no feedback closed",
			&valve{
				Tag:          "WWG-XV003",
				Description:  "Test valve 3",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV003_FBO",
				FbcTag:       "",
				FboAddress:   "I0.1",
				FbcAddress:   "",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       false,
			},
			map[string]string{
				"Tag":          "WWG-XV003",
				"Description":  "Test valve 3",
				"IDB":          "IDB_WWG-XV003",
				"Output":       `"WWG-XV003"`,
				"FBO":          `"WWG-XV003_FBO"`,
				"FBC":          `NOT "IDB_WWG-XV003".Q_On`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback open address",
			&valve{
				Tag:          "WWG-XV002",
				Description:  "Test valve 2",
				ActAddress:   "Q0.1",
				FboTag:       "WWG-XV002_FBO",
				FbcTag:       "WWG-XV002_FBC",
				FboAddress:   "M0.1",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV002",
				"Description":  "Test valve 2",
				"IDB":          "IDB_WWG-XV002",
				"Output":       `"WWG-XV002"`,
				"FBO":          `"WWG-XV002_FBO"`,
				"FBC":          `"WWG-XV002_FBC"`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback closed address",
			&valve{
				Tag:          "WWG-XV003",
				Description:  "Test valve 3",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV003_FBO",
				FbcTag:       "WWG-XV003_FBC",
				FboAddress:   "I0.1",
				FbcAddress:   "M0.2",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV003",
				"Description":  "Test valve 3",
				"IDB":          "IDB_WWG-XV003",
				"Output":       `"WWG-XV003"`,
				"FBO":          `"WWG-XV003_FBO"`,
				"FBC":          `"WWG-XV003_FBC"`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no output address",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.0",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I0.1",
				FbcAddress:   "I0.2",
				MonTimeOpen:  10,
				MonTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV001",
				"Description":  "Test valve 1",
				"IDB":          "IDB_WWG-XV001",
				"Output":       `"WWG-XV001"`,
				"FBO":          `"WWG-XV001_FBO"`,
				"FBC":          `"WWG-XV001_FBC"`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.InputMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valve.InputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valve_PlcTags(t *testing.T) {
	tests := []struct {
		name  string
		v     *valve
		wantT []*PlcTag
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I1.0",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       true,
			},
			[]*PlcTag{
				{Name: "WWG-XV001", Dtype: "Bool", Address: "Q0.2", Comment: "Test valve 1"},
				{Name: "WWG-XV001_FBO", Dtype: "Bool", Address: "I1.0", Comment: "Test valve 1 feedback open"},
				{Name: "WWG-XV001_FBC", Dtype: "Bool", Address: "I1.1", Comment: "Test valve 1 feedback closed"},
			},
		},
		{
			"Valve no feedback open",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			[]*PlcTag{
				{Name: "WWG-XV001", Dtype: "Bool", Address: "Q0.2", Comment: "Test valve 1"},
				{Name: "WWG-XV001_FBC", Dtype: "Bool", Address: "I1.1", Comment: "Test valve 1 feedback closed"},
			},
		},
		{
			"Valve no feedback closed",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "",
				FboAddress:   "I1.0",
				FbcAddress:   "",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       false,
			},
			[]*PlcTag{
				{Name: "WWG-XV001", Dtype: "Bool", Address: "Q0.2", Comment: "Test valve 1"},
				{Name: "WWG-XV001_FBO", Dtype: "Bool", Address: "I1.0", Comment: "Test valve 1 feedback open"},
			},
		},
		{
			"Valve no feedbacks",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "",
				FbcTag:       "",
				FboAddress:   "",
				FbcAddress:   "",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       false,
				hasFbc:       false,
			},
			[]*PlcTag{
				{Name: "WWG-XV001", Dtype: "Bool", Address: "Q0.2", Comment: "Test valve 1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := tt.v.PlcTags(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("valve.PlcTags() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_valve_outputPlcTag(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want *PlcTag
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I1.0",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       true,
			},
			&PlcTag{Name: "WWG-XV001", Dtype: "Bool", Address: "Q0.2", Comment: "Test valve 1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.outputPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valve.outputPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valve_fboPlcTag(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want *PlcTag
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I1.0",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       true,
			},
			&PlcTag{Name: "WWG-XV001_FBO", Dtype: "Bool", Address: "I1.0", Comment: "Test valve 1 feedback open"},
		},
		{
			"Valve no feedback open",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I1.0",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.fboPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valve.fboPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valve_fbcPlcTag(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want *PlcTag
	}{
		{
			"Valve",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "WWG-XV001_FBC",
				FboAddress:   "I1.0",
				FbcAddress:   "I1.1",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       true,
			},
			&PlcTag{Name: "WWG-XV001_FBC", Dtype: "Bool", Address: "I1.1", Comment: "Test valve 1 feedback closed"},
		},
		{
			"Valve no feedback closed",
			&valve{
				Tag:          "WWG-XV001",
				Description:  "Test valve 1",
				ActAddress:   "Q0.2",
				FboTag:       "WWG-XV001_FBO",
				FbcTag:       "",
				FboAddress:   "I1.0",
				FbcAddress:   "",
				MonTimeOpen:  10,
				MonTimeClose: 10,
				hasFbo:       true,
				hasFbc:       false,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.fbcPlcTag(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valve.fbcPlcTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
