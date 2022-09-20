package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type measmon struct {
	Tag         string
	Description string
	Unit        string
	Address     string
	Direct      bool
	LowLimit    float64
	HighLimit   float64
}

func NewMeasmon(tag, description, unit, address, direct, lowLimit, highLimit string) (*measmon, error) {
	directBool, err := strconv.ParseBool(direct)
	if err != nil {
		return nil, err
	}
	lowLimitFloat, err := strconv.ParseFloat(lowLimit, 64)
	if err != nil {
		return nil, err
	}
	highLimitFloat, err := strconv.ParseFloat(highLimit, 64)
	if err != nil {
		return nil, err
	}

	m := &measmon{
		Tag:         tag,
		Description: description,
		Unit:        unit,
		Address:     address,
		Direct:      directBool,
		LowLimit:    lowLimitFloat,
		HighLimit:   highLimitFloat,
	}

	if len(m.Address) == 0 {
		m.Address = "MW0"
		utils.Sugar.Warnw("No input address given",
			"measmon", m.Tag,
			"default", m.Address)
	}

	if m.LowLimit >= m.HighLimit {
		m.LowLimit = 0.0
		m.HighLimit = 100.0
		utils.Sugar.Warnw(
			"Low limit must be higher than high limit",
			"measmon", m.Tag,
			"low limit", lowLimitFloat,
			"high limit", highLimitFloat,
			"defaults", []float64{m.LowLimit, m.HighLimit},
		)
	}

	utils.Sugar.Infow("Object created",
		"measmon", m)

	return m, nil
}

func (m *measmon) InputMap() map[string]string {
	return map[string]string{
		"Tag":         m.Tag,
		"Description": m.Description,
		"IDB":         "IDB_" + m.Tag,
		"Unit":        m.Unit,
		"Input":       strconv.Quote(m.Tag),
		"LowLimit":    strconv.FormatFloat(m.LowLimit, 'f', 1, 64),
		"HighLimit":   strconv.FormatFloat(m.HighLimit, 'f', 1, 64),
	}
}

func (m *measmon) PlcTags() (t []*PlcTag) {
	t = append(t, &PlcTag{
		Name:    m.Tag,
		Dtype:   "Int",
		Address: m.Address,
		Comment: m.Description,
	})
	return
}
