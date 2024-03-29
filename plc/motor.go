package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/logger"
)

// TODO Separate tag name from output address tag name
// Make sure that double tags are not created in the tag tables

type motor struct {
	Tag              string
	Description      string
	ContactorAddress string
	FeedbackTag      string
	FeedbackAddress  string
	BreakerTag       string
	BreakerAddress   string
	SwitchTag        string
	SwitchAddress    string
	hasFeedback      bool
	hasBreaker       bool
	hasSwitch        bool
	Data             map[string]string
}

func NewMotor(tag, description, contactorAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, switchTag, switchAddress string, data map[string]string) (*motor, error) {
	m := &motor{
		Tag:              tag,
		Description:      description,
		ContactorAddress: contactorAddress,
		FeedbackTag:      feedbackTag,
		FeedbackAddress:  feedbackAddress,
		BreakerTag:       breakerTag,
		BreakerAddress:   breakerAddress,
		SwitchTag:        switchTag,
		SwitchAddress:    switchAddress,
		hasFeedback:      len(feedbackTag) > 0,
		hasBreaker:       len(breakerTag) > 0,
		hasSwitch:        len(switchTag) > 0,
		Data:             data,
	}

	if len(m.ContactorAddress) == 0 {
		m.ContactorAddress = "M0.0"
		logger.Sugar.Infow("No contactor address given",
			"motor", m.Tag,
			"default", m.ContactorAddress,
		)
	}
	if m.hasFeedback && len(m.FeedbackAddress) == 0 {
		m.FeedbackAddress = "M0.1"
		logger.Sugar.Infow("No feedback address given",
			"motor", m.Tag,
			"default", m.FeedbackAddress,
		)
	}
	if m.hasBreaker && len(m.BreakerAddress) == 0 {
		m.BreakerAddress = "M0.2"
		logger.Sugar.Infow("No breaker address given",
			"motor", m.Tag,
			"default", m.BreakerAddress,
		)
	}
	if m.hasSwitch && len(m.SwitchAddress) == 0 {
		m.SwitchAddress = "M0.3"
		logger.Sugar.Infow("No switch address given",
			"motor", m.Tag,
			"default", m.SwitchAddress,
		)
	}

	logger.Sugar.Debugw("Object created",
		"motor", m,
	)

	return m, nil
}

func (m *motor) InputMap() map[string]string {
	var feedbackTag, breakerTag, switchTag string

	if m.hasFeedback {
		feedbackTag = strconv.Quote(m.feedbackPlcTag().Name)
	} else {
		feedbackTag = strconv.Quote("IDB_"+m.Tag) + ".Q_On"
	}
	if m.hasBreaker {
		breakerTag = strconv.Quote(m.breakerPlcTag().Name)
	} else {
		breakerTag = "true"
	}
	if m.hasSwitch {
		switchTag = strconv.Quote(m.switchPlcTag().Name)
	} else {
		switchTag = "false"
	}

	// TODO: Add monitoring time
	input := map[string]string{
		"Tag":          m.Tag,
		"Description":  m.Description,
		"IDB":          "IDB_" + m.Tag,
		"ContactorTag": strconv.Quote(m.contactorPlcTag().Name),
		"FeedbackTag":  feedbackTag,
		"BreakerTag":   breakerTag,
		"SwitchTag":    switchTag,
	}
	for k, v := range m.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"motor", m.Tag,
				"name", k,
				"data", v)
		}
	}
	return input
}

func (m *motor) PlcTags() (t []*PlcTag) {
	if p := m.contactorPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := m.feedbackPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := m.breakerPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := m.switchPlcTag(); p != nil {
		t = append(t, p)
	}

	return
}

func (m *motor) contactorPlcTag() *PlcTag {
	return &PlcTag{
		Name:    m.Tag,
		Dtype:   "Bool",
		Address: m.ContactorAddress,
		Comment: m.Description,
	}
}

func (m *motor) feedbackPlcTag() *PlcTag {
	if !m.hasFeedback {
		return nil
	}
	return &PlcTag{
		Name:    m.FeedbackTag,
		Dtype:   "Bool",
		Address: m.FeedbackAddress,
		Comment: m.Description + " feedback",
	}
}

func (m *motor) breakerPlcTag() *PlcTag {
	if !m.hasBreaker {
		return nil
	}
	return &PlcTag{
		Name:    m.BreakerTag,
		Dtype:   "Bool",
		Address: m.BreakerAddress,
		Comment: m.Description + " breaker",
	}
}

func (m *motor) switchPlcTag() *PlcTag {
	if !m.hasSwitch {
		return nil
	}
	return &PlcTag{
		Name:    m.SwitchTag,
		Dtype:   "Bool",
		Address: m.SwitchAddress,
		Comment: m.Description + " protection switch",
	}
}
