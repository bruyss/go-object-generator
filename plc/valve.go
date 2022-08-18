package plc

import (
	"encoding/json"
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
)

type valve struct {
	Tag          string
	Description  string
	ActAddress   string
	FboTag       string
	FbcTag       string
	FboAddress   string
	FbcAddress   string
	MonTimeOpen  int
	MonTimeClose int
	hasFbo       bool
	hasFbc       bool
}

func NewValve(tag, description, actAddress, fboTag, fbcTag, fboAddress, fbcAddress, monTimeOpen, monTimeClose string) (*valve, error) {
	monTimeOpenInt, err := strconv.Atoi(monTimeOpen)
	if err != nil {
		return nil, err
	}
	monTimeCloseInt, err := strconv.Atoi(monTimeClose)
	if err != nil {
		return nil, err
	}
	v := &valve{
		Tag:          tag,
		Description:  description,
		ActAddress:   actAddress,
		FboTag:       fboTag,
		FbcTag:       fbcTag,
		FboAddress:   fboAddress,
		FbcAddress:   fbcAddress,
		MonTimeOpen:  monTimeOpenInt,
		MonTimeClose: monTimeCloseInt,
		hasFbo:       len(fboTag) > 0,
		hasFbc:       len(fbcTag) > 0,
	}

	if len(v.ActAddress) == 0 {
		v.ActAddress = "M0.0"
	}

	if v.hasFbo && len(v.FboAddress) == 0 {
		v.FboAddress = "M0.1"
	}

	if v.hasFbc && len(v.FbcAddress) == 0 {
		v.FbcAddress = "M0.2"
	}

	utils.Sugar.Debug("Object created",
		"Valve", v.String(),
	)
	return v, nil
}

func (v *valve) String() string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (v *valve) InputMap() map[string]string {
	var fbo, fbc string
	if v.hasFbo {
		fbo = v.FboTag
	} else {
		fbo = strconv.Quote("IDB_"+v.Tag) + ".Q_On"
	}
	if v.hasFbc {
		fbc = v.FbcTag
	} else {
		fbc = "NOT " + strconv.Quote("IDB_"+v.Tag) + ".Q_On"
	}

	return map[string]string{
		"Tag":          v.Tag,
		"Description":  v.Description,
		"IDB":          "IDB_" + v.Tag,
		"Output":       v.Tag,
		"FBO":          fbo,
		"FBC":          fbc,
		"MonTimeOpen":  strconv.Itoa(v.MonTimeOpen),
		"MonTimeClose": strconv.Itoa(v.MonTimeClose),
	}
}

func (v *valve) PlcTags() (t []*PlcTag) {
	t = append(t, v.outputPlcTag())

	if p := v.fboPlcTag(); p != nil {
		t = append(t, p)
	}

	if p := v.fbcPlcTag(); p != nil {
		t = append(t, p)
	}

	return
}

func (v *valve) outputPlcTag() *PlcTag {
	return &PlcTag{
		name:    v.Tag,
		dtype:   "Bool",
		address: v.ActAddress,
		comment: v.Description,
	}
}

func (v *valve) fboPlcTag() *PlcTag {
	if !v.hasFbo {
		return nil
	}
	return &PlcTag{
		name:    v.FboTag,
		dtype:   "Bool",
		address: v.FboAddress,
		comment: v.Description + " feedback open",
	}
}

func (v *valve) fbcPlcTag() *PlcTag {
	if !v.hasFbc {
		return nil
	}
	return &PlcTag{
		name:    v.FbcTag,
		dtype:   "Bool",
		address: v.FbcAddress,
		comment: v.Description + " feedback closed",
	}
}
