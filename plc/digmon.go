package plc

import "strconv"

type Digmon struct {
	tag         string
	description string
	address     string
	invert      bool
	alarm       bool
	invertAlarm bool
}

func (d *Digmon) InputMap() map[string]string {
	return map[string]string{
		"tag":         d.tag,
		"description": d.description,
		"idb":         "IDB_" + d.tag,
		"input":       d.tag,
		"invert":      strconv.FormatBool(d.invert),
		"alarm":       strconv.FormatBool(d.alarm),
		"invertalarm": strconv.FormatBool(d.invertAlarm),
	}
}

func (d *Digmon) PlcTags() []PlcTag {
	return []PlcTag{}
}
