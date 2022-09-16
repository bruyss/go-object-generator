package plc

type PlcTag struct {
	Name    string
	Dtype   string
	Address string
	Comment string
}

func (t *PlcTag) String() string {
	return t.Name + "(" + t.Dtype + ";" + t.Address + ";" + t.Comment + ")"
}
