package plc

import (
	"strconv"
	"strings"

	"github.com/bruyss/go-object-generator/logger"
)

// A digmon contains the information needed to create a digmon in the PLC project
type digmon struct {
	// Tag is the tag name of the object
	Tag string
	// Description is a short description of the functionality of the object.
	Description string
	// Address is the IO address of the digital input of the digmon
	Address string
	// Invert reports if the input should be inverted for logic
	Invert bool
	// Alarm reports if the input in the wrong state should cause an alarm
	Alarm bool
	// InvertAlarm inverts the logic of the alarm
	InvertAlarm bool
	// Data is any additional data that can be used in the object generation.
	Data map[string]string
}

// NewDigmon returns a reference to a new digmon. An error is returned if the entered information cannot be
// used to create a new object.
//
// The data used for the new object is a combination of the input data and sensible defaults if data is
// missing.
func NewDigmon(
	tag, description, address, invert, alarm, invertAlarm string,
	data map[string]string,
) (*digmon, error) {

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
		logger.Sugar.Infow("No input address given",
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
