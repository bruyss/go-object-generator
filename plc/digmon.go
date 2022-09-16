package plc

import (
	"strconv"
	"strings"

	"github.com/bruyss/go-object-generator/utils"
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
		utils.Sugar.Debugw("No input address given",
			"Digmon", d.Tag,
		)
		d.Address = "M0.0"
	}

	utils.Sugar.Debugw("Object created",
		"Digmon", d.String(),
	)

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
		"Invert":      strings.ToUpper(strconv.FormatBool(d.Invert)),
		"Alarm":       strings.ToUpper(strconv.FormatBool(d.Alarm)),
		"InvertAlarm": strings.ToUpper(strconv.FormatBool(d.InvertAlarm)),
	}
}

func (d *digmon) PlcTags() []*PlcTag {
	return []*PlcTag{
		{Name: d.Tag, Dtype: "Bool", Address: d.Address, Comment: d.Description},
	}
}
