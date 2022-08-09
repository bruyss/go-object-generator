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

func (m *Measmon) Stringer() string {
	return m.Tag
}

func (m *Measmon) InputMap() map[string]string {
	return map[string]string{
		"Tag":         m.Tag,
		"Description": m.Description,
		"IDB":         "IDB_" + m.Tag,
		"Unit":        m.Unit,
		"Input":       m.Tag,
		"LowLimit":    strconv.FormatFloat(m.LowLimit, 'f', 1, 64),
		"HighLimit":   strconv.FormatFloat(m.HighLimit, 'f', 1, 64),
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
