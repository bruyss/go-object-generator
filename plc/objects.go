package plc

type PLCObject interface {
	Stringer() string
	InputMap() map[string]string
	PlcTags() []PlcTag
}
