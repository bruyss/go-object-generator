package plc

import (
	"encoding/json"
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type freqMotor struct {
	Tag              string
	Description      string
	ContactorAddress string
	PqwAddress       string
	FeedbackTag      string
	FeedbackAddress  string
	BreakerTag       string
	BreakerAddress   string
	SwitchTag        string
	SwitchAddress    string
	DanfossDrive     bool
	hasFeedback      bool
	hasBreaker       bool
	hasSwitch        bool
}

func NewFreqMotor(tag, description, contactorAddress, pqwAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, switchTag, switchAddress, danfossDrive string) (*freqMotor, error) {
	danfossDriveBool, err := strconv.ParseBool(danfossDrive)
	if err != nil {
		return nil, err
	}

	f := &freqMotor{
		Tag:              tag,
		Description:      description,
		ContactorAddress: contactorAddress,
		PqwAddress:       pqwAddress,
		FeedbackTag:      feedbackTag,
		FeedbackAddress:  feedbackAddress,
		BreakerTag:       breakerTag,
		BreakerAddress:   breakerAddress,
		SwitchTag:        switchTag,
		SwitchAddress:    switchAddress,
		DanfossDrive:     danfossDriveBool,
		hasFeedback:      len(feedbackTag) > 0,
		hasBreaker:       len(breakerTag) > 0,
		hasSwitch:        len(switchTag) > 0,
	}

	if len(f.ContactorAddress) == 0 {
		f.ContactorAddress = "M0.0"
	}
	if len(f.PqwAddress) == 0 {
		f.PqwAddress = "MW0"
	}
	if f.hasFeedback && len(f.FeedbackAddress) == 0 {
		f.FeedbackAddress = "M0.1"
	}
	if f.hasBreaker && len(f.BreakerAddress) == 0 {
		f.BreakerAddress = "M0.2"
	}
	if f.hasSwitch && len(f.SwitchAddress) == 0 {
		f.SwitchAddress = "M0.3"
	}

	utils.Sugar.Debug("Object created",
		"freqMotor", f.String(),
	)

	return f, nil
}

func (f *freqMotor) String() string {
	b, _ := json.Marshal(f)
	return string(b)
}

func (f *freqMotor) InputMap() map[string]string {
	return map[string]string{
		"Tag":              f.Tag,
		"Description":      f.Description,
		"IDB":              "IDB_" + f.Tag,
		"ContactorAddress": f.ContactorAddress,
		"PQWAddress":       f.PqwAddress,
		"FeedbackTag":      f.FeedbackTag,
		"FeedbackAddress":  f.FeedbackAddress,
		"BreakerTag":       f.BreakerTag,
		"BreakerAddress":   f.BreakerAddress,
		"SwitchTag":        f.SwitchTag,
		"SwitchAddress":    f.SwitchAddress,
		"Danfoss":          strconv.FormatBool(f.DanfossDrive),
	}
}

func (f *freqMotor) PlcTags() (t []*PlcTag) {

	if p := f.contactorPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := f.pqwPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := f.feedbackPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := f.breakerPlcTag(); p != nil {
		t = append(t, p)
	}
	if p := f.switchPlcTag(); p != nil {
		t = append(t, p)
	}

	return
}

func (f *freqMotor) contactorPlcTag() *PlcTag {
	if f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		name:    f.Tag,
		dtype:   "Bool",
		address: f.ContactorAddress,
		comment: f.Description,
	}
}

func (f *freqMotor) pqwPlcTag() *PlcTag {
	if f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		name:    f.Tag + "_PQW",
		dtype:   "Int",
		address: f.PqwAddress,
		comment: f.Description + " output",
	}
}

func (f *freqMotor) feedbackPlcTag() *PlcTag {
	if !f.hasFeedback || f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		name:    f.FeedbackTag,
		dtype:   "Bool",
		address: f.FeedbackAddress,
		comment: f.Description + " feedback",
	}
}

func (f *freqMotor) breakerPlcTag() *PlcTag {
	if !f.hasBreaker {
		return nil
	}
	return &PlcTag{
		name:    f.BreakerTag,
		dtype:   "Bool",
		address: f.BreakerAddress,
		comment: f.Description + " breaker",
	}
}

func (f *freqMotor) switchPlcTag() *PlcTag {
	if !f.hasSwitch {
		return nil
	}
	return &PlcTag{
		name:    f.SwitchTag,
		dtype:   "Bool",
		address: f.SwitchAddress,
		comment: f.Description + " protection switch",
	}
}
