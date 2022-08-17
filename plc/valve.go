package plc

import (
	"strconv"

	"github.com/bruyss/go-object-generator/utils"
	"go.uber.org/zap"
)

type valve struct {
	tag          string
	description  string
	actAddress   string
	fboTag       string
	fbcTag       string
	fboAddress   string
	fbcAddress   string
	monTimeOpen  int
	monTimeClose int
	hasFbo       bool
	hasFbc       bool
}

func NewValve(tag, description, actAddress, fboTag, fbcTag, fboAddress, fbcAddress string, monTimeOpen, monTimeClose int) (v *valve) {
	v = &valve{
		tag:          tag,
		description:  description,
		actAddress:   actAddress,
		fboTag:       fboTag,
		fbcTag:       fbcTag,
		fboAddress:   fboAddress,
		fbcAddress:   fbcAddress,
		monTimeOpen:  monTimeOpen,
		monTimeClose: monTimeClose,
		hasFbo:       len(fboTag) > 0,
		hasFbc:       len(fbcTag) > 0,
	}

	if len(v.actAddress) == 0 {
		v.actAddress = "M0.0"
	}

	if v.hasFbo && len(v.fboAddress) == 0 {
		v.fboAddress = "M0.1"
	}

	if v.hasFbc && len(v.fbcAddress) == 0 {
		v.fbcAddress = "M0.2"
	}

	utils.Sugar.Debug(
		"Valve created",
		zap.String("Valve", v.tag),
	)
	return
}

func (v *valve) Tag() string {
	return v.tag
}

func (v *valve) InputMap() map[string]string {
	var fbo, fbc string
	if v.hasFbo {
		fbo = v.fboTag
	} else {
		fbo = strconv.Quote("IDB_"+v.tag) + ".Q_On"
	}
	if v.hasFbc {
		fbc = v.fbcTag
	} else {
		fbc = "NOT " + strconv.Quote("IDB_"+v.tag) + ".Q_On"
	}

	return map[string]string{
		"Tag":          v.tag,
		"Description":  v.description,
		"IDB":          "IDB_" + v.tag,
		"Output":       v.tag,
		"FBO":          fbo,
		"FBC":          fbc,
		"MonTimeOpen":  strconv.Itoa(v.monTimeOpen),
		"MonTimeClose": strconv.Itoa(v.monTimeClose),
	}
}

func (v *valve) PlcTags() (t []*PlcTag) {
	t = append(t, &PlcTag{
		name:    v.tag,
		dtype:   "Bool",
		address: v.actAddress,
		comment: v.description,
	})

	if v.hasFbo {
		t = append(t, &PlcTag{
			name:    v.fboTag,
			dtype:   "Bool",
			address: v.fboAddress,
			comment: v.description + " feedback open",
		})
	}

	if v.hasFbc {
		t = append(t, &PlcTag{
			name:    v.fbcTag,
			dtype:   "Bool",
			address: v.fbcAddress,
			comment: v.description + " feedback closed",
		})
	}
	return
}
