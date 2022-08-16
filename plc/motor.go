package plc

type motor struct {
	tag         string
	description string
}

func (m *motor) Tag() string {
	return m.tag
}

func (m *motor) InputMap() map[string]string {
	return make(map[string]string)
}

func (m *motor) PlcTags() (t []*PlcTag) {
	return
}
