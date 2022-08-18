package plc

import (
	"strconv"
)

type digmon struct {
	Tag         string
	Description string
	Address     string
	Invert      bool
	Alarm       bool
	InvertAlarm bool
}

func NewDigmon(tag, description, address, invert, alarm, invertAlarm string) (*digmon, error) {

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
	}

	if len(d.Address) == 0 {
		d.Address = "M0.0"
	}

	return d, nil
}

func (d *digmon) String() string {
	return d.Tag
}

func (d *digmon) InputMap() map[string]string {
	return map[string]string{
		"Tag":         d.Tag,
		"Description": d.Description,
		"IDB":         "IDB_" + d.Tag,
		"Input":       strconv.Quote(d.Tag),
		"Invert":      strconv.FormatBool(d.Invert),
		"Alarm":       strconv.FormatBool(d.Alarm),
		"InvertAlarm": strconv.FormatBool(d.InvertAlarm),
	}
}

func (d *digmon) PlcTags() []*PlcTag {
	return []*PlcTag{
		{name: d.Tag, dtype: "Bool", address: d.Address, comment: d.Description},
	}
}
