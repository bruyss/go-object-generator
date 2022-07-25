package plc

type PID struct {
	tag         string
	description string
}

func (p *PID) InputMap() map[string]string {
	return make(map[string]string)
}

func (p *PID) PlcTags() []PlcTag {
	return []PlcTag{}
}
