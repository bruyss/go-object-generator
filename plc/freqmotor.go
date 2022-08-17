package plc

import "strconv"

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
}

func NewFreqMotor(tag, description, contactorAddress, pqwAddress, feedbackTag, feedbackAddress, breakerTag, breakerAddress, switchTag, switchAddress string, danfossDrive bool) *freqMotor {
	return &freqMotor{}
}

func (f *freqMotor) Stringer() string {
	return f.Tag
}

func (f *freqMotor) InputMap() map[string]string {
	// var feedbackAddress, breakerAddress, switchAddress string
	// if len(f.feedbackTag) > 0 && len(f.feedbackAddress) == 0 {
	// 	feedbackAddress = "M0.0"
	// } else {
	// 	feedbackAddress = f.feedbackAddress
	// }
	return map[string]string{
		"Tag":              f.Tag,
		"Description":      f.Description,
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

func (f *freqMotor) PlcTags() []PlcTag {
	return []PlcTag{}
}
