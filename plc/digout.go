package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/logger"
)

// TODO Implement digital output

type digout struct {
	Tag             string
	Description     string
	OutputAddress   string
	FeedbackTag     string
	FeedbackAddress string
	BreakerTag      string
	BreakerAddress  string
	hasFeedback     bool
	hasBreaker      bool
	MonitoringTime  int
	Data            map[string]string
}

// NewDigout returns a pointer to a digital out object with the given information
func NewDigout(tag, description, outputAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, monitoringTime string, data map[string]string) (*digout, error) {
	// Create object reference
	d := &digout{
		Tag:         tag,
		Description: description,
		Data:        data,
	}

	// Monitoring time
	monitoringTimeInt, err := strconv.Atoi(monitoringTime)
	if err != nil {
		d.MonitoringTime = 10
		logger.Sugar.Warnf("Monitoring time cannot be parsed",
			"digout", d.Tag,
			"monitoring time", monitoringTime,
			"default", d.MonitoringTime,
			"error", err,
		)
	} else {
		d.MonitoringTime = monitoringTimeInt
	}

	// Output address
	if len(outputAddress) > 0 {
		d.OutputAddress = outputAddress
	} else {
		d.OutputAddress = "M0.0"
		logger.Sugar.Infow("No output address given",
			"digout", d.Tag,
			"default", d.OutputAddress,
		)
	}

	// Feedback
	if len(feedbackTag) > 0 {
		d.FeedbackTag = feedbackTag
		d.hasFeedback = true
		if len(feedbackAddress) > 0 {
			d.FeedbackAddress = feedbackAddress
		} else {
			d.FeedbackAddress = "M0.1"
			logger.Sugar.Infow("No feedback address given",
				"digout", d.Tag,
				"default", d.FeedbackAddress,
			)
		}
	} else {
		d.hasFeedback = false
	}

	// Breaker
	if len(breakerTag) > 0 {
		d.BreakerTag = breakerTag
		d.hasBreaker = true
		if len(breakerAddress) > 0 {
			d.BreakerAddress = breakerAddress
		} else {
			d.BreakerAddress = "M0.2"
			logger.Sugar.Infow("No breaker address given",
				"digout", d.Tag,
				"default", d.BreakerAddress,
			)
		}
	} else {
		d.hasBreaker = false
	}

	logger.Sugar.Debugw("Object created",
		"digout", d)

	return d, nil
}

func (d *digout) InputMap() map[string]string {
	var feedbackTag, breakerTag string

	if d.hasFeedback {
		feedbackTag = strconv.Quote(d.feedbackTag().Name)
	} else {
		feedbackTag = strconv.Quote("IDB_"+d.Tag) + ".Q_On"
	}

	if d.hasBreaker {
		breakerTag = strconv.Quote(d.breakerTag().Name)
	} else {
		breakerTag = "true"
	}

	input := map[string]string{
		"Tag":            d.Tag,
		"Description":    d.Description,
		"IDB":            "IDB_" + d.Tag,
		"OutputTag":      strconv.Quote(d.outputTag().Name),
		"FeedbackTag":    feedbackTag,
		"BreakerTag":     breakerTag,
		"MonitoringTime": strconv.Itoa(d.MonitoringTime),
	}

	for k, v := range d.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"digout", d.Tag,
				"name", k,
				"data", v)
		}
	}

	return input
}

func (d *digout) PlcTags() (t []*PlcTag) {
	if p := d.outputTag(); p != nil {
		t = append(t, p)
	}
	if p := d.feedbackTag(); p != nil {
		t = append(t, p)
	}
	if p := d.breakerTag(); p != nil {
		t = append(t, p)
	}

	return
}

func (d *digout) outputTag() *PlcTag {
	return &PlcTag{
		Name:    d.Tag,
		Dtype:   "Bool",
		Address: d.OutputAddress,
		Comment: d.Description,
	}
}

func (d *digout) feedbackTag() *PlcTag {
	if !d.hasFeedback {
		return nil
	}

	return &PlcTag{
		Name:    d.FeedbackTag,
		Dtype:   "Bool",
		Address: d.FeedbackAddress,
		Comment: d.Description + " feedback",
	}
}

func (d *digout) breakerTag() *PlcTag {
	if !d.hasBreaker {
		return nil
	}

	return &PlcTag{
		Name:    d.BreakerTag,
		Dtype:   "Bool",
		Address: d.BreakerAddress,
		Comment: d.Description + " breaker",
	}
}
