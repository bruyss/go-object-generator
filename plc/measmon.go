package plc

import "strconv"

type Measmon struct {
	Tag         string
	Description string
	Unit        string
	Address     string
	Direct      bool
	LowLimit    float64
	HighLimit   float64
}

func (m *Measmon) InputMap() map[string]string {
	return map[string]string{
		"tag":         m.Tag,
		"description": m.Description,
		"idb":         "IDB_" + m.Tag,
		"unit":        m.Unit,
		"input":       m.Tag,
		"lowlim":      strconv.FormatFloat(m.LowLimit, 'f', 1, 64),
		"highlim":     strconv.FormatFloat(m.HighLimit, 'f', 1, 64),
	}
}

func (m *Measmon) PlcTags() []PlcTag {
	inputTag := PlcTag{
		name:    m.Tag,
		dtype:   "Int",
		address: m.Address,
		comment: m.Description,
	}
	return []PlcTag{inputTag}
}
