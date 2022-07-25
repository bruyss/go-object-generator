package plc

type Valve struct {
	tag          string
	description  string
	actAddress   string
	fboTag       string
	fbcTag       string
	fboAddress   string
	fbcAddress   string
	monTimeOpen  int
	monTimeClose int
}

func (v *Valve) InputMap() map[string]string {
	return make(map[string]string)
}

func (v *Valve) PlcTags() []PlcTag {
	return []PlcTag{}
}
