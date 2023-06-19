package config

import (
	"github.com/spf13/viper"
)

// SetDefaults defines the default viper settings for the tool
func SetDefaults() {
	// Generation settings
	viper.SetDefault("gensettings.general.secondpulse", "iSecPulse")
	viper.SetDefault("gensettings.general.simulation", "iSimulation")
	viper.SetDefault("gensettings.general.todobit", "\"DB_Test\".TODO_BOOL")
	viper.SetDefault("gensettings.general.todoreal", "\"DB_Test\".TODO_REAL")
	viper.SetDefault("gensettings.general.wincc", true)

	viper.SetDefault("gensettings.measmon.objectfb", "FB_Measmon")
	viper.SetDefault("gensettings.measmon.callfc", "Measmon_Call")
	viper.SetDefault("gensettings.measmon.hmidb", "HMI_Measmons")
	viper.SetDefault("gensettings.measmon.hmitype", "HMI_Measmon")
	viper.SetDefault("gensettings.measmon.startindex", 0)
	viper.SetDefault("gensettings.measmon.tagtable", "Measmons")

	viper.SetDefault("gensettings.digmon.objectfb", "FB_Digmon")
	viper.SetDefault("gensettings.digmon.callfc", "Digmon_Call")
	viper.SetDefault("gensettings.digmon.hmidb", "HMI_Digmons")
	viper.SetDefault("gensettings.digmon.hmitype", "HMI_Digmon")
	viper.SetDefault("gensettings.digmon.startindex", 0)
	viper.SetDefault("gensettings.digmon.tagtable", "Digmons")

	viper.SetDefault("gensettings.valve.objectfb", "FB_Valve")
	viper.SetDefault("gensettings.valve.callfc", "Valve_Call")
	viper.SetDefault("gensettings.valve.hmidb", "HMI_Valves")
	viper.SetDefault("gensettings.valve.hmitype", "HMI_Valve")
	viper.SetDefault("gensettings.valve.startindex", 0)
	viper.SetDefault("gensettings.valve.tagtable", "Valves")

	viper.SetDefault("gensettings.controlvalve.objectfb", "FB_ControlValve")
	viper.SetDefault("gensettings.controlvalve.callfc", "ControlValve_Call")
	viper.SetDefault("gensettings.controlvalve.hmidb", "HMI_ControlValves")
	viper.SetDefault("gensettings.controlvalve.hmitype", "HMI_ControlValve")
	viper.SetDefault("gensettings.controlvalve.startindex", 0)
	viper.SetDefault("gensettings.controlvalve.tagtable", "ControlValves")

	viper.SetDefault("gensettings.motor.objectfb", "FB_Motor")
	viper.SetDefault("gensettings.motor.callfc", "Motor_Call")
	viper.SetDefault("gensettings.motor.hmidb", "HMI_Motors")
	viper.SetDefault("gensettings.motor.hmitype", "HMI_Motor")
	viper.SetDefault("gensettings.motor.startindex", 0)
	viper.SetDefault("gensettings.motor.tagtable", "Motors")

	viper.SetDefault("gensettings.digout.objectfb", "FB_DigitalOut")
	viper.SetDefault("gensettings.digout.callfc", "DigitalOut_Call")
	viper.SetDefault("gensettings.digout.hmidb", "HMI_DigitalOut")
	viper.SetDefault("gensettings.digout.hmitype", "HMI_DigitalOut")
	viper.SetDefault("gensettings.digout.startindex", 0)
	viper.SetDefault("gensettings.digout.tagtable", "DigitalOuts")

	viper.SetDefault("gensettings.freqmotor.objectfb", "FB_Motor_Freq")
	viper.SetDefault("gensettings.freqmotor.callfc", "FreqMotor_Call")
	viper.SetDefault("gensettings.freqmotor.hmidb", "HMI_Motor_Freq")
	viper.SetDefault("gensettings.freqmotor.hmitype", "HMI_FreqMotor")
	viper.SetDefault("gensettings.freqmotor.startindex", 0)
	viper.SetDefault("gensettings.freqmotor.tagtable", "FreqMotors")

	// File names
	viper.SetDefault("filenames.general.objectsource", "excelsource_go.xlsx")
	viper.SetDefault("filenames.general.idbtemplate", "idb.tmpl")
	viper.SetDefault("filenames.general.hmidbtemplate", "hmidb.tmpl")

	viper.SetDefault("filenames.controlvalve.sourcetemplate", "controlValve.tmpl")
	viper.SetDefault("filenames.controlvalve.idbfile", "ControlValve_IDBs.db")
	viper.SetDefault("filenames.controlvalve.hmidbfile", "ControlValve_HMIDB.db")
	viper.SetDefault("filenames.controlvalve.sourcefile", "ControlValve_source.scl")
	viper.SetDefault("filenames.controlvalve.tagfile", "ControlValve_tags.xml")

	viper.SetDefault("filenames.digmon.sourcetemplate", "digmon.tmpl")
	viper.SetDefault("filenames.digmon.idbfile", "Digmon_IDBs.db")
	viper.SetDefault("filenames.digmon.hmidbfile", "Digmon_HMIDB.db")
	viper.SetDefault("filenames.digmon.sourcefile", "Digmon_source.scl")
	viper.SetDefault("filenames.digmon.tagfile", "Digmon_tags.xml")

	viper.SetDefault("filenames.freqmotor.sourcetemplate", "freqMotor.tmpl")
	viper.SetDefault("filenames.freqmotor.idbfile", "FreqMotor_IDBs.db")
	viper.SetDefault("filenames.freqmotor.hmidbfile", "FreqMotor_HMIDB.db")
	viper.SetDefault("filenames.freqmotor.sourcefile", "FreqMotor_source.scl")
	viper.SetDefault("filenames.freqmotor.tagfile", "FreqMotor_tags.xml")

	viper.SetDefault("filenames.measmon.sourcetemplate", "measmon.tmpl")
	viper.SetDefault("filenames.measmon.idbfile", "Measmon_IDBs.db")
	viper.SetDefault("filenames.measmon.hmidbfile", "Measmon_HMIDB.db")
	viper.SetDefault("filenames.measmon.sourcefile", "Measmon_source.scl")
	viper.SetDefault("filenames.measmon.tagfile", "Measmon_tags.xml")

	viper.SetDefault("filenames.motor.sourcetemplate", "motor.tmpl")
	viper.SetDefault("filenames.motor.idbfile", "Motor_IDBs.db")
	viper.SetDefault("filenames.motor.hmidbfile", "Motor_HMIDB.db")
	viper.SetDefault("filenames.motor.sourcefile", "Motor_source.scl")
	viper.SetDefault("filenames.motor.tagfile", "Motor_tags.xml")

	viper.SetDefault("filenames.digout.sourcetemplate", "digout.tmpl")
	viper.SetDefault("filenames.digout.idbfile", "DigitalOut_IDBs.db")
	viper.SetDefault("filenames.digout.hmidbfile", "DigitalOut_HMIDB.db")
	viper.SetDefault("filenames.digout.sourcefile", "DigitalOut_source.scl")
	viper.SetDefault("filenames.digout.tagfile", "DigitalOut_tags.xml")

	viper.SetDefault("filenames.valve.sourcetemplate", "valve.tmpl")
	viper.SetDefault("filenames.valve.idbfile", "Valve_IDBs.db")
	viper.SetDefault("filenames.valve.hmidbfile", "Valve_HMIDB.db")
	viper.SetDefault("filenames.valve.sourcefile", "Valve_source.scl")
	viper.SetDefault("filenames.valve.tagfile", "Valve_tags.xml")
}
