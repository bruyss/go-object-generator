package plc

type ControlValve struct {
	tag         string
	description string
}

func (c *ControlValve) InputMap() map[string]string {
	return make(map[string]string)
}

func (c *ControlValve) PlcTags() []PlcTag {
	return []PlcTag{}
}
