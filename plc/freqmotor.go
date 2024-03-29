package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/logger"
)

// TODO Separate tag name from output address tag name
// Make sure that double tags are not created in the tag tables

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
	AlarmTag         string
	AlarmAddress     string
	DanfossDrive     bool
	hasFeedback      bool
	hasBreaker       bool
	hasSwitch        bool
	hasAlarm         bool
	Data             map[string]string
}

func NewFreqMotor(tag, description, contactorAddress, pqwAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, switchTag, switchAddress, alarmTag, alarmAddress, danfossDrive string, data map[string]string) (*freqMotor, error) {
	var danfossDriveBool bool
	if danfossDrive == "" {
		danfossDriveBool = false
	} else {
		var err error
		danfossDriveBool, err = strconv.ParseBool(danfossDrive)
		if err != nil {
			return nil, err
		}
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
		AlarmTag:         alarmTag,
		AlarmAddress:     alarmAddress,
		DanfossDrive:     danfossDriveBool,
		hasFeedback:      len(feedbackTag) > 0,
		hasBreaker:       len(breakerTag) > 0,
		hasSwitch:        len(switchTag) > 0,
		hasAlarm:         len(alarmTag) > 0,
		Data:             data,
	}

	if len(f.ContactorAddress) == 0 && !f.DanfossDrive {
		f.ContactorAddress = "M0.0"
		logger.Sugar.Infow("No contactor address given",
			"frequency motor", f.Tag,
			"default", f.ContactorAddress,
		)
	}
	if len(f.PqwAddress) == 0 && !f.DanfossDrive {
		f.PqwAddress = "MW0"
		logger.Sugar.Infow("No PQW address given",
			"frequency motor", f.Tag,
			"default", f.PqwAddress,
		)
	}
	if f.hasFeedback && len(f.FeedbackAddress) == 0 && !f.DanfossDrive {
		f.FeedbackAddress = "M0.1"
		logger.Sugar.Infow("No feedback address given",
			"frequency motor", f.Tag,
			"default", f.FeedbackAddress,
		)
	}
	if f.hasBreaker && len(f.BreakerAddress) == 0 {
		f.BreakerAddress = "M0.2"
		logger.Sugar.Infow("No breaker address given",
			"frequency motor", f.Tag,
			"default", f.BreakerAddress,
		)
	}
	if f.hasSwitch && len(f.SwitchAddress) == 0 {
		f.SwitchAddress = "M0.3"
		logger.Sugar.Infow("No switch address given",
			"frequency motor", f.Tag,
			"default", f.SwitchAddress,
		)
	}
	if f.hasAlarm && len(f.AlarmAddress) == 0 && !f.DanfossDrive {
		f.AlarmAddress = "M0.4"
		logger.Sugar.Infow("No alarm address given",
			"frequency motor", f.Tag,
			"default", f.AlarmAddress,
		)
	}

	logger.Sugar.Debugw("Object created",
		"freq motor", f,
	)

	return f, nil
}

func (f *freqMotor) InputMap() map[string]string {
	var feedbackTag, breakerTag, switchTag string
	if f.hasFeedback {
		feedbackTag = strconv.Quote(f.FeedbackTag)
	} else {
		feedbackTag = strconv.Quote("IDB_"+f.Tag) + ".Q_On"
	}
	if f.hasBreaker {
		breakerTag = strconv.Quote(f.BreakerTag)
	} else {
		breakerTag = "FALSE"
	}
	if f.hasSwitch {
		switchTag = strconv.Quote(f.SwitchTag)
	} else {
		switchTag = "TRUE"
	}

	input := map[string]string{
		"Tag":          f.Tag,
		"Description":  f.Description,
		"IDB":          "IDB_" + f.Tag,
		"ContactorTag": strconv.Quote(f.Tag),
		"PQW":          strconv.Quote(f.Tag + "_PQW"),
		"FeedbackTag":  feedbackTag,
		"BreakerTag":   breakerTag,
		"SwitchTag":    switchTag,
		"AlarmTag":     strconv.Quote(f.Tag + "_AL"),
		"Danfoss":      strconv.FormatBool(f.DanfossDrive),
	}
	for k, v := range f.Data {
		_, exists := input[k]
		if !exists {
			input[k] = v
			logger.Sugar.Debugw("Additional data added to input map",
				"freq motor", f.Tag,
				"name", k,
				"data", v)
		}
	}
	return input
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
	if p := f.alarmPlcTag(); p != nil {
		t = append(t, p)
	}
	_ = f.alarmPlcTag()

	return
}

func (f *freqMotor) contactorPlcTag() *PlcTag {
	if f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		Name:    f.Tag,
		Dtype:   "Bool",
		Address: f.ContactorAddress,
		Comment: f.Description,
	}
}

func (f *freqMotor) pqwPlcTag() *PlcTag {
	if f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		Name:    f.Tag + "_PQW",
		Dtype:   "Int",
		Address: f.PqwAddress,
		Comment: f.Description + " output",
	}
}

func (f *freqMotor) feedbackPlcTag() *PlcTag {
	if !f.hasFeedback || f.DanfossDrive {
		return nil
	}
	return &PlcTag{
		Name:    f.FeedbackTag,
		Dtype:   "Bool",
		Address: f.FeedbackAddress,
		Comment: f.Description + " feedback",
	}
}

func (f *freqMotor) breakerPlcTag() *PlcTag {
	if !f.hasBreaker {
		return nil
	}
	return &PlcTag{
		Name:    f.BreakerTag,
		Dtype:   "Bool",
		Address: f.BreakerAddress,
		Comment: f.Description + " breaker",
	}
}

func (f *freqMotor) switchPlcTag() *PlcTag {
	if !f.hasSwitch {
		return nil
	}
	return &PlcTag{
		Name:    f.SwitchTag,
		Dtype:   "Bool",
		Address: f.SwitchAddress,
		Comment: f.Description + " protection switch",
	}
}

func (f *freqMotor) alarmPlcTag() *PlcTag {
	if !f.hasAlarm {
		return nil
	}
	return &PlcTag{
		Name:    f.AlarmTag,
		Dtype:   "Bool",
		Address: f.AlarmAddress,
		Comment: f.Description + " drive alarm",
	}
}
