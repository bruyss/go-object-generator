package plc

type PLCObject interface {
	InputMap() map[string]string
	PlcTags() []PlcTag
}
