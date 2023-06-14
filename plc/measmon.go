package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/logger"
)

type measmon struct {
	Tag         string
	Description string
	Unit        string
	Address     string
	Direct      bool
	LowLimit    float64
	HighLimit   float64
	Data        map[string]string
}

func NewMeasmon(tag, description, unit, address, direct, lowLimit, highLimit string, data map[string]string) (*measmon, error) {
	directBool, err := strconv.ParseBool(direct)
	if err != nil {
		return nil, err
	}
	lowLimitFloat, err := strconv.ParseFloat(lowLimit, 64)
	if err != nil {
		lowLimitFloat = 0.0
		logger.Sugar.Warnw("Cannot parse low limit to float",
			"measmon", tag,
			"value", lowLimit,
			"default", lowLimitFloat)
	}
	highLimitFloat, err := strconv.ParseFloat(highLimit, 64)
	if err != nil {
		highLimitFloat = 100.0
		logger.Sugar.Warnw("Cannot parse high limit to float",
			"measmon", tag,
			"value", highLimit,
			"default", highLimitFloat)
	}

	m := &measmon{
		Tag:         tag,
		Description: description,
		Unit:        unit,
		Address:     address,
		Direct:      directBool,
		LowLimit:    lowLimitFloat,
		HighLimit:   highLimitFloat,
		Data:        data,
	}

	if len(m.Unit) == 0 {
		logger.Sugar.Infow("No unit given",
			"measmon", m.Tag)
	}

	if len(m.Address) == 0 {
		m.Address = "MW0"
		logger.Sugar.Infow("No input address given",
			"measmon", m.Tag,
			"default", m.Address)
	}

	if m.LowLimit >= m.HighLimit {
		m.LowLimit = 0.0
		m.HighLimit = 100.0
		logger.Sugar.Infow(
			"Low limit must be higher than high limit",
			"measmon", m.Tag,
			"low limit", lowLimitFloat,
			"high limit", highLimitFloat,
			"defaults", []float64{m.LowLimit, m.HighLimit},
		)
	}

	logger.Sugar.Debugw("Object created",
		"measmon", m)

	return m, nil
}

func (m *measmon) InputMap() map[string]string {
	input := map[string]string{
		"Tag":         m.Tag,
		"Description": m.Description,
		"IDB":         "IDB_" + m.Tag,
		"Unit":        m.Unit,
		"Input":       strconv.Quote(m.Tag),
		"LowLimit":    strconv.FormatFloat(m.LowLimit, 'f', 1, 64),
		"HighLimit":   strconv.FormatFloat(m.HighLimit, 'f', 1, 64),
	}
	for k, v := range m.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"measmon", m.Tag,
				"name", k,
				"data", v)
		}
	}
	return input
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
