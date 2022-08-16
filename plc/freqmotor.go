package plc

type freqMotor struct {
	tag         string
	description string
}

func (f *freqMotor) Tag() string {
	return f.tag
}

func (f *freqMotor) InputMap() map[string]string {
	return make(map[string]string)
}

func (f *freqMotor) PlcTags() (t []*PlcTag) {
	return
}
