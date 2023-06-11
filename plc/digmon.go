package plc

import (
	"strconv"
	"strings"

	"github.com/bruyss/go-object-generator/logger"
)

type digmon struct {
	Tag         string
	Description string
	Address     string
	Invert      bool
	Alarm       bool
	InvertAlarm bool
	Data        map[string]string
}

func NewDigmon(tag, description, address, invert, alarm, invertAlarm string, data map[string]string) (*digmon, error) {

	invertBool, err := strconv.ParseBool(invert)
	if err != nil {
		return nil, err
	}
	alarmBool, err := strconv.ParseBool(alarm)
	if err != nil {
		return nil, err
	}
	invertAlarmBool, err := strconv.ParseBool(invertAlarm)
	if err != nil {
		return nil, err
	}

	d := &digmon{
		Tag:         tag,
		Description: description,
		Address:     address,
		Invert:      invertBool,
		Alarm:       alarmBool,
		InvertAlarm: invertAlarmBool,
		Data:        data,
	}

	if len(d.Address) == 0 {
		d.Address = "M0.0"
		logger.Sugar.Warnw("No input address given",
			"digmon", d.Tag,
			"default", d.Address,
		)
	}

	logger.Sugar.Debugw("Object created",
		"Digmon", d,
	)

	return d, nil
}

func (d *digmon) InputMap() map[string]string {
	input := map[string]string{
		"Tag":         d.Tag,
		"Description": d.Description,
		"IDB":         "IDB_" + d.Tag,
		"Input":       strconv.Quote(d.Tag),
		"Invert":      strings.ToUpper(strconv.FormatBool(d.Invert)),
		"Alarm":       strings.ToUpper(strconv.FormatBool(d.Alarm)),
		"InvertAlarm": strings.ToUpper(strconv.FormatBool(d.InvertAlarm)),
	}
	for k, v := range d.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"digmon", d.Tag,
				"name", k,
				"data", v)
		}
	}
	return input
}

func (d *digmon) PlcTags() []*PlcTag {
	return []*PlcTag{
		{Name: d.Tag, Dtype: "Bool", Address: d.Address, Comment: d.Description},
	}
}
