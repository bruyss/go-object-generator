package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
	"go.uber.org/zap"
)

type measmon struct {
	tag         string
	description string
	unit        string
	address     string
	direct      bool
	lowLimit    float64
	highLimit   float64
}

func NewMeasmon(tag, description, unit, address, direct, lowLimit, highLimit string) *measmon {
	directBool, _ := strconv.ParseBool(direct)
	lowLimitFloat, _ := strconv.ParseFloat(lowLimit, 64)
	highLimitFloat, _ := strconv.ParseFloat(highLimit, 64)

	if lowLimit >= highLimit {
		utils.Sugar.Info(
			"Low limit must be higher than high limit",
			zap.String("tag", tag),
			zap.Float64("lowLimit", lowLimitFloat),
			zap.Float64("highLimit", highLimitFloat),
		)
		lowLimitFloat = 0.0
		highLimitFloat = 100.0
	}

	return &measmon{
		tag:         tag,
		description: description,
		unit:        unit,
		address:     address,
		direct:      directBool,
		lowLimit:    lowLimitFloat,
		highLimit:   highLimitFloat,
	}
}

func (m *measmon) Tag() string {
	return m.tag
}

func (m *measmon) InputMap() map[string]string {
	return map[string]string{
		"Tag":         m.tag,
		"Description": m.description,
		"IDB":         "IDB_" + m.tag,
		"Unit":        m.unit,
		"Input":       m.tag,
		"LowLimit":    strconv.FormatFloat(m.lowLimit, 'f', 1, 64),
		"HighLimit":   strconv.FormatFloat(m.highLimit, 'f', 1, 64),
	}
}

func (m *measmon) PlcTags() (t []*PlcTag) {
	t = append(t, &PlcTag{
		name:    m.tag,
		dtype:   "Int",
		address: m.address,
		comment: m.description,
	})
	return
}
