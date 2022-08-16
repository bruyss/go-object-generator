package plc

type PLCObject interface {
	String() string
	InputMap() map[string]string
	PlcTags() []PlcTag
}
