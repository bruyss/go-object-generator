package obwriter

import (
	"encoding/json"
	"os"

	"github.com/bruyss/go-object-generator/utils"
)

type generalSettings struct {
	SecondPulse string
	Simulation  string
	TodoBit     string
	TodoReal    string
	Wincc       bool
}

type objectSettings struct {
	ObjectFb   string
	CallFc     string
	HmiDb      string
	StartIndex int
	TagTable   string
}

type generatorSettings struct {
	General      *generalSettings
	Measmon      *objectSettings
	Digmon       *objectSettings
	Valve        *objectSettings
	ControlValve *objectSettings
	Motor        *objectSettings
	FreqMotor    *objectSettings
}

var defaultSettings = &generatorSettings{
	General: &generalSettings{
		SecondPulse: "iSecPulse",
		Simulation:  "iSimulation",
		TodoBit:     `"DB_Test".TODO_BIT`,
		TodoReal:    `"DB_Test".TODO_REAL`,
		Wincc:       false,
	},
	Measmon: &objectSettings{
		ObjectFb:   "FB_Measmon",
		CallFc:     "Measmons_Call",
		HmiDb:      "HMI_Measmons",
		StartIndex: 0,
		TagTable:   "Measmons",
	},
	Digmon: &objectSettings{
		ObjectFb:   "FB_Digmon",
		CallFc:     "Digmons_Call",
		HmiDb:      "HMI_Digmons",
		StartIndex: 0,
		TagTable:   "Digmons",
	},
	Valve: &objectSettings{
		ObjectFb:   "FB_Valve",
		CallFc:     "Valves_Call",
		HmiDb:      "HMI_Valves",
		StartIndex: 0,
		TagTable:   "Valves",
	},
	ControlValve: &objectSettings{
		ObjectFb:   "FB_ControlValve",
		CallFc:     "ControlValves_Call",
		HmiDb:      "HMI_ControlValves",
		StartIndex: 0,
		TagTable:   "ControlValves",
	},
	Motor: &objectSettings{
		ObjectFb:   "FB_Motor",
		CallFc:     "Motors_Call",
		HmiDb:      "HMI_Motors",
		StartIndex: 0,
		TagTable:   "Motors",
	},
	FreqMotor: &objectSettings{
		ObjectFb:   "FB_Motor_Freq",
		CallFc:     "FreqMotors_Call",
		HmiDb:      "HMI_FreqMotors",
		StartIndex: 0,
		TagTable:   "FreqMotors",
	},
}

func (s *generatorSettings) writeSettings(name string) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		utils.Sugar.Error(err)
	}

	err = os.WriteFile(name, b, 0777)
	if err != nil {
		utils.Sugar.Error(err)
	}
}

func WriteDefaultSettings(name string) {
	defaultSettings.writeSettings(name)
}

func ReadSettings(name string) *generatorSettings {
	b, err := os.ReadFile(name)
	if err != nil {
		utils.Sugar.Error(err)
	}
	var s *generatorSettings
	err = json.Unmarshal(b, &s)
	if err != nil {
		utils.Sugar.Error(err)
	}
	return s
}
