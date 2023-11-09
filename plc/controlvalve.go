package plc

import (
	"strconv"
	"strings"

	"github.com/bruyss/go-object-generator/logger"
)

// A controlValve contains the information needed to create a control valve in the PLC project
type controlValve struct {
	// Tag is the tag name of the object
	Tag string
	// Description is a short description of the functionality of the object.
	Description string
	// OutputTag is the tag name given to the analog output of the control valve.
	OutputTag string
	// OutputAddress is the IO address given to the analog output of the control valve.
	OutputAddress string
	// FeedbackTag is the tag name given to the analog input giving position feedback of the control valve.
	FeedbackTag string
	// FeedbackAddress is the IO address given to the analog input giving position feedback of the control
	// valve.
	FeedbackAddress string
	// MonitoringTime is the delay before a difference in output and feedback will cause an object alarm.
	MonitoringTime int
	// hasFeedback is a flag to indicate if a feedback input is present.
	hasFeedback bool
	// Data is any additional data that can be used in the object generation.
	Data map[string]string
}

// NewControlValve returns a reference to a new control valve. An error is returned if the entered
// information cannot be used to create a new object.
//
// The data used for the new object is a combination of the input data and sensible defaults if data is
// missing.
func NewControlValve(
	tag, description, outputTag, outputAddress, feedbackTag, feedbackAddress, monitoringTime string,
	data map[string]string,
) (*controlValve, error) {
	monitoringTimeInt, err := strconv.Atoi(monitoringTime)
	if err != nil {
		return nil, err
	}

	c := &controlValve{
		Tag:             tag,
		Description:     description,
		OutputTag:       outputTag,
		OutputAddress:   outputAddress,
		FeedbackTag:     feedbackTag,
		FeedbackAddress: feedbackAddress,
		MonitoringTime:  monitoringTimeInt,
		hasFeedback:     len(feedbackTag) > 0,
		Data:            data,
	}

	if len(c.OutputTag) == 0 {
		c.OutputTag = tag
		logger.Sugar.Debugw("No output tag given, using default",
			"control valve", c.Tag,
			"default", c.OutputTag)
	}

	if len(c.OutputAddress) == 0 {
		c.OutputAddress = "MW0"
		logger.Sugar.Infow("No output address provided",
			"control valve", c.Tag,
			"default", c.OutputAddress)
	}

	if len(c.FeedbackAddress) == 0 && c.hasFeedback {
		c.FeedbackAddress = "MW2"
		logger.Sugar.Infow("No feedback address provided",
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
	input := map[string]string{
		"Tag":            c.Tag,
		"Description":    c.Description,
		"IDB":            "IDB_" + c.Tag,
		"NoFeedback":     strings.ToUpper(strconv.FormatBool(!c.hasFeedback)),
		"Feedback":       feedbackTag,
		"MonitoringTime": strconv.Itoa(c.MonitoringTime),
		"Output":         strconv.Quote(c.OutputTag),
	}
	for k, v := range c.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"control valve", c.Tag,
				"name", k,
				"data", v)
		}
	}
	return input
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
		Name:    c.OutputTag,
		Dtype:   "Int",
		Address: c.OutputAddress,
		Comment: c.Description + " output",
	}
}

func (c *controlValve) feedbackPlcTag() *PlcTag {
	if !c.hasFeedback {
		return nil
	}
	return &PlcTag{
		Name:    c.FeedbackTag,
		Dtype:   "Int",
		Address: c.FeedbackAddress,
		Comment: c.Description + " feedback",
	}
}
