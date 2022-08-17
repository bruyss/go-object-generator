package plc

type PlcTag struct {
	name    string
	dtype   string
	address string
	comment string
}

func (t *PlcTag) String() string {
	return t.name + "(" + t.dtype + ";" + t.address + ";" + t.comment + ")"
}
