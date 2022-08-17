package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type digmon struct {
	tag         string
	description string
	address     string
	invert      bool
	alarm       bool
	invertAlarm bool
}

func NewDigmon(tag, description, address string, invert, alarm, invertAlarm bool) *digmon {
	if len(address) <= 0 {
		address = "M0.0"
	}
	return &digmon{
		tag:         tag,
		description: description,
		address:     address,
		invert:      invert,
		alarm:       alarm,
		invertAlarm: invertAlarm,
	}
}

func (d *digmon) Tag() string {
	return d.tag
}

func (d *digmon) InputMap() map[string]string {
	return map[string]string{
		"Tag":         d.tag,
		"Description": d.description,
		"IDB":         "IDB_" + d.tag,
		"Input":       utils.TagQuotes(d.tag),
		"Invert":      strconv.FormatBool(d.invert),
		"Alarm":       strconv.FormatBool(d.alarm),
		"InvertAlarm": strconv.FormatBool(d.invertAlarm),
	}
}

func (d *digmon) PlcTags() []*PlcTag {
	return []*PlcTag{
		{name: d.tag, dtype: "Bool", address: d.address, comment: d.description},
	}
}
