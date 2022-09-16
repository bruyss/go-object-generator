package plc

import (
	"encoding/json"
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

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
}

func NewMotor(tag, description, contactorAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, switchTag, switchAddress string) (*motor, error) {
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
	}

	if len(m.ContactorAddress) == 0 {
		m.ContactorAddress = "M0.0"
	}
	if m.hasFeedback && len(m.FeedbackAddress) == 0 {
		m.FeedbackAddress = "M0.1"
	}
	if m.hasBreaker && len(m.BreakerAddress) == 0 {
		m.BreakerAddress = "M0.2"
	}
	if m.hasSwitch && len(m.SwitchAddress) == 0 {
		m.SwitchAddress = "M0.3"
	}

	utils.Sugar.Debug("Object created",
		"motor", m.String(),
	)

	return m, nil
}

func (m *motor) String() string {
	b, _ := json.Marshal(m)
	return string(b)
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
	// TODO: Add contactor tag
	return map[string]string{
		"Tag":          m.Tag,
		"Description":  m.Description,
		"IDB":          strconv.Quote("IDB_" + m.Tag),
		"ContactorTag": strconv.Quote(m.contactorPlcTag().Name),
		"FeedbackTag":  feedbackTag,
		"BreakerTag":   breakerTag,
		"SwitchTag":    switchTag,
	}
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
