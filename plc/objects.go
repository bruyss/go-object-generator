package plc

type PlcObject interface {
	Tag() string
	InputMap() map[string]string
	PlcTags() []*PlcTag
}
