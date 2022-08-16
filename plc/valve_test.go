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
		monTimeOpen  int
		monTimeClose int
	}
	tests := []struct {
		name  string
		args  args
		wantV *valve
	}{
		{
			"Valve",
			args{"WWG-XV001", "Test valve 1", "Q0.0", "WWG-XV001_FBO", "WWG-XV001_FBC", "I0.1", "I0.2", 10, 15},
			&valve{
				tag:          "WWG-XV001",
				description:  "Test valve 1",
				actAddress:   "Q0.0",
				fboTag:       "WWG-XV001_FBO",
				fbcTag:       "WWG-XV001_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
		},
		{
			"Valve no feedback open",
			args{"WWG-XV002", "Test valve 2", "Q0.1", "", "WWG-XV002_FBC", "", "I0.2", 10, 10},
			&valve{
				tag:          "WWG-XV002",
				description:  "Test valve 2",
				actAddress:   "Q0.1",
				fboTag:       "",
				fbcTag:       "WWG-XV002_FBC",
				fboAddress:   "",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
		},
		{
			"Valve no feedback closed",
			args{"WWG-XV003", "Test valve 3", "Q0.2", "WWG-XV003_FBO", "", "I0.1", "", 10, 15},
			&valve{
				tag:          "WWG-XV003",
				description:  "Test valve 3",
				actAddress:   "Q0.2",
				fboTag:       "WWG-XV003_FBO",
				fbcTag:       "",
				fboAddress:   "I0.1",
				fbcAddress:   "",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       false,
			},
		},
		{
			"Valve no feedback open address",
			args{"WWG-XV002", "Test valve 2", "Q0.1", "WWG-XV002_FBO", "WWG-XV002_FBC", "", "I0.2", 10, 10},
			&valve{
				tag:          "WWG-XV002",
				description:  "Test valve 2",
				actAddress:   "Q0.1",
				fboTag:       "WWG-XV002_FBO",
				fbcTag:       "WWG-XV002_FBC",
				fboAddress:   "M0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 10,
				hasFbo:       true,
				hasFbc:       true,
			},
		},
		{
			"Valve no feedback closed address",
			args{"WWG-XV003", "Test valve 3", "Q0.2", "WWG-XV003_FBO", "WWG-XV003_FBC", "I0.1", "", 10, 15},
			&valve{
				tag:          "WWG-XV003",
				description:  "Test valve 3",
				actAddress:   "Q0.2",
				fboTag:       "WWG-XV003_FBO",
				fbcTag:       "WWG-XV003_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "M0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
		},
		{
			"Valve no output address",
			args{"WWG-XV001", "Test valve 1", "", "WWG-XV001_FBO", "WWG-XV001_FBC", "I0.1", "I0.2", 10, 15},
			&valve{
				tag:          "WWG-XV001",
				description:  "Test valve 1",
				actAddress:   "M0.0",
				fboTag:       "WWG-XV001_FBO",
				fbcTag:       "WWG-XV001_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := NewValve(tt.args.tag, tt.args.description, tt.args.actAddress, tt.args.fboTag, tt.args.fbcTag, tt.args.fboAddress, tt.args.fbcAddress, tt.args.monTimeOpen, tt.args.monTimeClose); !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("NewValve() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func Test_valve_Tag(t *testing.T) {
	tests := []struct {
		name string
		v    *valve
		want string
	}{
		{
			"Valve",
			&valve{
				tag:          "WWG-XV001",
				description:  "Test valve 1",
				actAddress:   "Q0.0",
				fboTag:       "WWG-XV001_FBO",
				fbcTag:       "WWG-XV001_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			"WWG-XV001",
		},
		{
			"Valve no feedback open",
			&valve{
				tag:          "WWG-XV002",
				description:  "Test valve 2",
				actAddress:   "Q0.1",
				fboTag:       "",
				fbcTag:       "WWG-XV002_FBC",
				fboAddress:   "",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			"WWG-XV002",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Tag(); got != tt.want {
				t.Errorf("valve.Tag() = %v, want %v", got, tt.want)
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
				tag:          "WWG-XV001",
				description:  "Test valve 1",
				actAddress:   "Q0.0",
				fboTag:       "WWG-XV001_FBO",
				fbcTag:       "WWG-XV001_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV001",
				"Description":  "Test valve 1",
				"IDB":          "IDB_WWG-XV001",
				"Output":       "WWG-XV001",
				"FBO":          "WWG-XV001_FBO",
				"FBC":          "WWG-XV001_FBC",
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback open",
			&valve{
				tag:          "WWG-XV002",
				description:  "Test valve 2",
				actAddress:   "Q0.1",
				fboTag:       "",
				fbcTag:       "WWG-XV002_FBC",
				fboAddress:   "",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 10,
				hasFbo:       false,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV002",
				"Description":  "Test valve 2",
				"IDB":          "IDB_WWG-XV002",
				"Output":       "WWG-XV002",
				"FBO":          `"IDB_WWG-XV002".Q_On`,
				"FBC":          "WWG-XV002_FBC",
				"MonTimeOpen":  "10",
				"MonTimeClose": "10",
			},
		},
		{
			"Valve no feedback closed",
			&valve{
				tag:          "WWG-XV003",
				description:  "Test valve 3",
				actAddress:   "Q0.2",
				fboTag:       "WWG-XV003_FBO",
				fbcTag:       "",
				fboAddress:   "I0.1",
				fbcAddress:   "",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       false,
			},
			map[string]string{
				"Tag":          "WWG-XV003",
				"Description":  "Test valve 3",
				"IDB":          "IDB_WWG-XV003",
				"Output":       "WWG-XV003",
				"FBO":          "WWG-XV003_FBO",
				"FBC":          `NOT "IDB_WWG-XV003".Q_On`,
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback open address",
			&valve{
				tag:          "WWG-XV002",
				description:  "Test valve 2",
				actAddress:   "Q0.1",
				fboTag:       "WWG-XV002_FBO",
				fbcTag:       "WWG-XV002_FBC",
				fboAddress:   "M0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV002",
				"Description":  "Test valve 2",
				"IDB":          "IDB_WWG-XV002",
				"Output":       "WWG-XV002",
				"FBO":          "WWG-XV002_FBO",
				"FBC":          "WWG-XV002_FBC",
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no feedback closed address",
			&valve{
				tag:          "WWG-XV003",
				description:  "Test valve 3",
				actAddress:   "Q0.2",
				fboTag:       "WWG-XV003_FBO",
				fbcTag:       "WWG-XV003_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "M0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV003",
				"Description":  "Test valve 3",
				"IDB":          "IDB_WWG-XV003",
				"Output":       "WWG-XV003",
				"FBO":          "WWG-XV003_FBO",
				"FBC":          "WWG-XV003_FBC",
				"MonTimeOpen":  "10",
				"MonTimeClose": "15",
			},
		},
		{
			"Valve no output address",
			&valve{
				tag:          "WWG-XV001",
				description:  "Test valve 1",
				actAddress:   "Q0.0",
				fboTag:       "WWG-XV001_FBO",
				fbcTag:       "WWG-XV001_FBC",
				fboAddress:   "I0.1",
				fbcAddress:   "I0.2",
				monTimeOpen:  10,
				monTimeClose: 15,
				hasFbo:       true,
				hasFbc:       true,
			},
			map[string]string{
				"Tag":          "WWG-XV001",
				"Description":  "Test valve 1",
				"IDB":          "IDB_WWG-XV001",
				"Output":       "WWG-XV001",
				"FBO":          "WWG-XV001_FBO",
				"FBC":          "WWG-XV001_FBC",
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
			NewValve("WWG-XV001", "Test valve 1", "Q0.2", "WWG-XV001_FBO", "WWG-XV001_FBC", "I1.0", "I1.1", 10, 10),
			[]*PlcTag{
				{name: "WWG-XV001", dtype: "Bool", address: "Q0.2", comment: "Test valve 1"},
				{name: "WWG-XV001_FBO", dtype: "Bool", address: "I1.0", comment: "Test valve 1 feedback open"},
				{name: "WWG-XV001_FBC", dtype: "Bool", address: "I1.1", comment: "Test valve 1 feedback closed"},
			},
		},
		{
			"Valve no feedback open",
			NewValve("WWG-XV001", "Test valve 1", "Q0.2", "", "WWG-XV001_FBC", "", "I1.2", 10, 10),
			[]*PlcTag{
				{name: "WWG-XV001", dtype: "Bool", address: "Q0.2", comment: "Test valve 1"},
				{name: "WWG-XV001_FBC", dtype: "Bool", address: "I1.2", comment: "Test valve 1 feedback closed"},
			},
		},
		{
			"Valve no feedback closed",
			NewValve("WWG-XV001", "Test valve 1", "Q0.2", "WWG-XV001_FBO", "", "I1.0", "", 10, 10),
			[]*PlcTag{
				{name: "WWG-XV001", dtype: "Bool", address: "Q0.2", comment: "Test valve 1"},
				{name: "WWG-XV001_FBO", dtype: "Bool", address: "I1.0", comment: "Test valve 1 feedback open"},
			},
		},
		{
			"Valve no feedbacks",
			NewValve("WWG-XV001", "Test valve 1", "Q0.2", "", "", "", "", 10, 10),
			[]*PlcTag{
				{name: "WWG-XV001", dtype: "Bool", address: "Q0.2", comment: "Test valve 1"},
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
