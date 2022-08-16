package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type ControlValve struct {
	Tag             string
	Description     string
	Address         string
	FeedbackAddress string
	MonitoringTime  int
}

func (c *ControlValve) String() string {
	return c.Tag
}

func (c *ControlValve) InputMap() map[string]string {
	return map[string]string{
		"Tag":            c.Tag,
		"Description":    c.Description,
		"IDB":            "IDB_" + c.Tag,
		"NoFeedback":     strconv.FormatBool(len(c.FeedbackAddress) <= 0),
		"Feedback":       c.Tag + "_FB",
		"MonitoringTime": strconv.Itoa(c.MonitoringTime),
		"Output":         utils.TagQuotes(c.Tag),
	}
}

func (c *ControlValve) PlcTags() []PlcTag {
	tags := make([]PlcTag, 0)
	var address, feedbackAddress string

	if len(c.Address) > 0 {
		address = c.Address
	} else {
		address = "MW0"
	}

	if len(c.FeedbackAddress) > 0 {
		feedbackAddress = c.FeedbackAddress
	} else {
		feedbackAddress = "MW2"
	}

	tags = append(tags, PlcTag{
		name:    c.Tag,
		dtype:   "Int",
		address: address,
		comment: c.Description + " output",
	})

	tags = append(tags, PlcTag{
		name:    c.Tag + "_FB",
		dtype:   "Int",
		address: feedbackAddress,
		comment: c.Description + " feedback",
	})

	return tags
}
