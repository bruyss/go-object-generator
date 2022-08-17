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
		Tag:         tag,
		Description: description,
		Unit:        unit,
		Address:     address,
		Direct:      directBool,
		LowLimit:    lowLimitFloat,
		HighLimit:   highLimitFloat,
	}
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
		"Input":       utils.TagQuotes(m.Tag),
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
