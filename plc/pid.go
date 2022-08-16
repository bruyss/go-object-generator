package plc

type pid struct {
	tag         string
	description string
}

func (p *pid) Tag() string {
	return p.tag
}

func (p *pid) InputMap() map[string]string {
	return make(map[string]string)
}

func (p *pid) PlcTags() (t []*PlcTag) {
	return
}
