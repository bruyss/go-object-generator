package plc

type Motor struct {
	tag         string
	description string
}

func (m *Motor) InputMap() map[string]string {
	return make(map[string]string)
}

func (m *Motor) PlcTags() []PlcTag {
	return []PlcTag{}
}
