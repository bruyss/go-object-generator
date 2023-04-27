package plc

import (
	"strconv"
	"strings"

	"github.com/bruyss/go-object-generator/logger"
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
		logger.Sugar.Warnw("No output address provided",
			"control valve", c.Tag,
			"default", c.Address)
	}

	if len(c.FeedbackAddress) == 0 && c.hasFeedback {
		c.FeedbackAddress = "MW2"
		logger.Sugar.Warnw("No feedback address provided",
			"control valve", c.Tag,
			"default", c.FeedbackAddress)
	}

	logger.Sugar.Debugw("Object created",
		"control valve", c)

	return c, nil
}

func (c *controlValve) InputMap() map[string]string {
	var feedbackTag string
	if c.hasFeedback {
		feedbackTag = strconv.Quote(c.FeedbackTag)
	} else {
		feedbackTag = strconv.Quote("IDB_"+c.Tag) + ".Q_On"
	}
	return map[string]string{
		"Tag":            c.Tag,
		"Description":    c.Description,
		"IDB":            "IDB_" + c.Tag,
		"NoFeedback":     strings.ToUpper(strconv.FormatBool(!c.hasFeedback)),
		"Feedback":       feedbackTag,
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
