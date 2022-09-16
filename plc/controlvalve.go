package plc

import (
	"encoding/json"
	"strconv"
)

type controlValve struct {
	Tag             string
	Description     string
	Address         string
	FeedbackTag     string
	FeedbackAddress string
	MonitoringTime  int
	hasFeedback     bool
}

func NewControlValve(tag, description, address, feedbackTag, feedbackAddress, monitoringTime string) (*controlValve, error) {
	monitoringTimeInt, err := strconv.Atoi(monitoringTime)
	if err != nil {
		return nil, err
	}

	c := &controlValve{
		Tag:             tag,
		Description:     description,
		Address:         address,
		FeedbackTag:     feedbackTag,
		FeedbackAddress: feedbackAddress,
		MonitoringTime:  monitoringTimeInt,
		hasFeedback:     len(feedbackTag) > 0,
	}

	if len(c.Address) == 0 {
		c.Address = "MW0"
	}

	if len(c.FeedbackAddress) == 0 && c.hasFeedback {
		c.FeedbackAddress = "MW2"
	}

	return c, nil
}

func (c *controlValve) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c *controlValve) InputMap() map[string]string {
	return map[string]string{
		"Tag":            c.Tag,
		"Description":    c.Description,
		"IDB":            "IDB_" + c.Tag,
		"NoFeedback":     strconv.FormatBool(!c.hasFeedback),
		"Feedback":       c.FeedbackTag,
		"MonitoringTime": strconv.Itoa(c.MonitoringTime),
		"Output":         strconv.Quote(c.outputPlcTag().Name),
	}
}

func (c *controlValve) PlcTags() (t []*PlcTag) {
	if p := c.outputPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := c.feedbackPlcTag(); p != nil {
		t = append(t, p)
	}
	return t
}

func (c *controlValve) outputPlcTag() *PlcTag {
	return &PlcTag{
		Name:    c.Tag,
		Dtype:   "Int",
		Address: c.Address,
		Comment: c.Description + " output",
	}
}

func (c *controlValve) feedbackPlcTag() *PlcTag {
	if !c.hasFeedback {
		return nil
	}
	return &PlcTag{
		Name:    c.Tag + "_FB",
		Dtype:   "Int",
		Address: c.FeedbackAddress,
		Comment: c.Description + " feedback",
	}
}
