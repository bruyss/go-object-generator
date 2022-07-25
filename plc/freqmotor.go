package plc

type FreqMotor struct {
	tag         string
	description string
}

func (f *FreqMotor) InputMap() map[string]string {
	return make(map[string]string)
}

func (f *FreqMotor) PlcTags() []PlcTag {
	return []PlcTag{}
}
