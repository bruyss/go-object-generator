// Package plc defines types which correspond to commonly used object types
// in PLC programming.
//
// The currently implemented PLC types are:
//  - Measmon: measuring and monitoring device
//  - Control valve: position controlled valve
//  - Digmon: digital monitoring device
//  - Digout: digital output
//  - FreqMotor: frequency controlled motor
//  - Motor: direct on line motor
//  - Valve: on-off valve
package plc

// A PlcObject contains information regarding one plc object instance
type PlcObject interface {
	// InputMap returns a map containing all information needed for object instance generation
	InputMap() map[string]string
	// PlcTags returns a list of PlcTags associated with the object instance
	PlcTags() []*PlcTag
}
