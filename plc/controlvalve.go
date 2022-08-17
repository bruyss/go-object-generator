package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type controlValve struct {
	tag             string
	description     string
	address         string
	feedbackAddress string
	monitoringTime  int
}

// TODO: write NewControlValve function with input checking, move out of PlcTags

func (c *controlValve) Tag() string {
	return c.tag
}

func (c *controlValve) InputMap() map[string]string {
	return map[string]string{
		"Tag":            c.tag,
		"Description":    c.description,
		"IDB":            "IDB_" + c.tag,
		"NoFeedback":     strconv.FormatBool(len(c.feedbackAddress) <= 0),
		"Feedback":       c.tag + "_FB",
		"MonitoringTime": strconv.Itoa(c.monitoringTime),
		"Output":         utils.TagQuotes(c.tag),
	}
}

func (c *controlValve) PlcTags() (t []*PlcTag) {
	var address, feedbackAddress string

	if len(c.address) > 0 {
		address = c.address
	} else {
		address = "MW0"
	}

	if len(c.feedbackAddress) > 0 {
		feedbackAddress = c.feedbackAddress
	} else {
		feedbackAddress = "MW2"
	}

	t = append(t, &PlcTag{
		name:    c.tag,
		dtype:   "Int",
		address: address,
		comment: c.description + " output",
	})

	t = append(t, &PlcTag{
		name:    c.tag + "_FB",
		dtype:   "Int",
		address: feedbackAddress,
		comment: c.description + " feedback",
	})

	return t
}
