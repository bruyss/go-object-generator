package plc

import (
	"encoding/json"
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
	"go.uber.org/zap"
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

	if m.LowLimit >= m.HighLimit {
		utils.Sugar.Info(
			"Low limit must be higher than high limit",
			zap.String("tag", tag),
			zap.Float64("lowLimit", lowLimitFloat),
			zap.Float64("highLimit", highLimitFloat),
		)
		m.LowLimit = 0.0
		m.HighLimit = 100.0
	}

	utils.Sugar.Debugw("Object created",
		"measmon", m.String())

	return m, nil
}

func (m *measmon) String() string {
	b, _ := json.Marshal(m)
	return string(b)
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
		name:    m.Tag,
		dtype:   "Int",
		address: m.Address,
		comment: m.Description,
	})
	return
}
