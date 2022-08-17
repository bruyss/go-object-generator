package plc

type PlcObject interface {
	// Tag() string
	String() string
	InputMap() map[string]string
	PlcTags() []*PlcTag
}
