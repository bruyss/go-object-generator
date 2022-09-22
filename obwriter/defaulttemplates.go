package obwriter

var Templates = map[string]string{
	"idbs":         idbTemplate,
	"hmidb":        hmidb,
	"controlValve": controlValveTemplate,
	"digmon":       digmonTemplate,
	"freqMotor":    freqMotorTemplate,
	"motor":        motorTemplate,
	"measmon":      measmonTemplate,
	"valve":        valveTemplate,
}

const (
	idbTemplate string = `{{$objectfb := .ObjectSettings.ObjectFb}}{{range .Objects}}DATA_BLOCK "{{.InputMap.IDB}}"
{ S7_Optimized_Access := 'TRUE' }
VERSION : 0.1
NON_RETAIN
{{$objectfb}}
BEGIN
END_DATA_BLOCK

{{end}}`

	hmidb string = `{{$os := .ObjectSettings -}}
DATA_BLOCK "{{.ObjectSettings.HmiDb}}"
{ S7_Optimized_Access := 'FALSE' }
VERSION : 0.1
NON_RETAIN
   VAR
      {{range .Objects -}}
      "{{.InputMap.Tag}}" : "{{$os.HmiType}}";
      {{end}}
   END_VAR

BEGIN

END_DATA_BLOCK`

	measmonTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap}}
    REGION {{$d.Tag}}: {{$d.Description}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                        Unit := '{{$d.Unit}}',
                        SecPuls := {{$gs.SecondPulse}},
                        Reset := TRUE,
                        Local := FALSE,
                        Simulation := {{$gs.Simulation}},
                        AnalogInput := {{$d.Input}},
                        LowLimit := {{$d.LowLimit}},
                        HighLimit := {{$d.HighLimit}},
                        {{if $gs.Wincc -}}
                            HMI := "{{$os.HmiDb}}"."{{$d.Tag}}"
                        {{- else -}}
                            HMI := "{{$os.HmiDb}}".o[{{$index}}]
                        {{- end}});
    END_REGION
{{end}}
END_FUNCTION`

	digmonTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap }}
    REGION {{$d.Tag}}: {{$d.Description}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                        SecPuls := {{$gs.SecondPulse}},
                        Reset := FALSE,
                        EnableAlarm := {{$d.Alarm}},
                        InvertAlarm := {{$d.InvertAlarm}},
                        Input := {{$d.Input}},
                        Invert := {{$d.Invert}},
                        {{if $gs.Wincc -}}
                            HMI := "{{$os.HmiDb}}"."{{$d.Tag}}"
                        {{- else -}}
                            HMI := "{{$os.HmiDb}}".o[{{$index}}]
                        {{- end}});
    END_REGION
{{end}}
END_FUNCTION`

	valveTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap}}
    REGION {{$d.Tag}}: {{$d.Description}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                        SecPuls := {{$gs.SecondPulse}},
                        Reset := FALSE,
                        Local := FALSE,
                        Simulation := {{$gs.Simulation}},
                        Permit := TRUE,
                        NO := FALSE,
                        Activate := {{$gs.TodoBit}},
                        FeedbackActivated := {{$d.FBO}},
                        FeedbackDeactivated := {{$d.FBC}},
                        MonitoringTimeAct := {{$d.MonTimeOpen}},
                        MonitoringTimeDeact := {{$d.MonTimeClose}},
                        Q_Activate := {{$d.Output}},
                        {{if $gs.Wincc -}}
                            HMI := "{{$os.HmiDb}}"."{{$d.Tag}}"
                        {{- else -}}
                            HMI := "{{$os.HmiDb}}".o[{{$index}}]
                        {{- end}});
    END_REGION
{{end}}
END_FUNCTION`

	controlValveTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap}}
    REGION {{$d.Tag}}: {{$d.Description}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                        Secpulse := {{$gs.SecondPulse}},
                        Reset := FALSE,
                        Local := FALSE,
                        Simulation := {{$gs.Simulation}},
                        Permit := TRUE,
                        NoFeedback := {{$d.NoFeedback}},
                        {{if eq $d.NoFeedback "false" -}}
                            Feedback := {{$d.Feedback}},
                        {{- end}}
                        SP := {{$gs.TodoReal}},
                        TimeMon := {{$d.MonitoringTime}},
                        PQW := {{$d.Output}},
                        {{if $gs.Wincc -}}
                            HMI := "{{$os.HmiDb}}"."{{$d.Tag}}"
                        {{- else -}}
                            HMI := "{{$os.HmiDb}}".o[{{$index}}]
                        {{- end}});
    END_REGION
{{end}}
END_FUNCTION`

	freqMotorTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
    iReset  : Bool;

END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings -}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap}}
    REGION {{$d.Tag}}: {{$d.Description}} {{$d.Danfoss}}
    {{- if eq $d.Danfoss "true"}}
        "FC_Danfoss"(iHW_ID := "{{$d.Tag}}~PPO_4_-_6_6_Words__Danfoss_Telegra,,,~PPO_4_-_6_6_Words__Danfoss_,,,",
                        iTagname := '{{$d.Tag}}',
                        iSecPulse := {{$gs.SecondPulse}},
                        iReset := iReset,
                        iSimulation := {{$gs.Simulation}},
                        iPermit := TRUE,
                        iActivate := {{$gs.TodoBit}},
                        iThermalProtection := {{$d.BreakerTag}},
                        iProtectionSwitch := {{$d.SwitchTag}},
                        iSetpoint := {{$gs.TodoReal}},
                        iMinimumFreq := 15.0,
                        iMaximumFreq := 50.0,
                        iMonitoringTime := 10,
                        {{if $gs.Wincc -}}
                            {{- "ioHMI"}} := "{{$os.HmiDb}}"."{{$d.Tag}}",
                        {{- else -}}
                            {{- "ioHMI"}} := "{{$os.HmiDb}}".o[{{$index}}],
                        {{- end}}
                        ioDrive := "DRIVE_COMM"."{{$d.Tag}}",
                        Motor_Freq_Instance := "{{$d.IDB}}");
    {{- else}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                        SecPuls := {{$gs.SecondPulse}},
                        Reset := TRUE,
                        Local := FALSE,
                        Simulation := {{$gs.Simulation}},
                        Permit := TRUE,
                        Activate := {{$gs.TodoBit}},
                        Feedback := {{$d.FeedbackTag}},
                        ThermalProt := {{$d.BreakerTag}},
                        Protectionswitch := {{$d.SwitchTag}},
                        DriveError := {{$d.AlarmTag}},
                        MonitoringTime := 10,
                        Setpoint := {{$gs.TodoReal}},
                        Q_On := {{$d.ContactorTag}},
                        PQW := {{$d.PQW}},
                        {{if $gs.Wincc -}}
                            {{- "HMI"}} := "{{$os.HmiDb}}"."{{$d.Tag}}");
                        {{- else -}}
                            {{- "HMI"}} := "{{$os.HmiDb}}".o[{{$index}}]);
                        {{- end}}
    {{- end}}
    END_REGION
{{end}}
END_FUNCTION`

	motorTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$gs := .GeneralSettings}}{{$os := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    {{- $d := $object.InputMap}}
    REGION {{$d.Tag}}: {{$d.Description}}
        "{{$d.IDB}}"(Tagname := '{{$d.Tag}}',
                            SecPuls := {{$gs.SecondPulse}},
                            Reset := TRUE,
                            Local := FALSE,
                            Simulation := {{$gs.Simulation}},
                            Permit := TRUE,
                            Activate := {{$gs.TodoBit}},
                            Feedback := {{$d.FeedbackTag}},
                            ThermalProt := {{$d.BreakerTag}},
                            Protectionswitch := {{$d.SwitchTag}},
                            MonitoringTime := 10,
                            Q_On := %M0.0,
                            {{if $gs.Wincc -}}
                                HMI := "{{$os.HmiDb}}"."{{$d.Tag}}"
                            {{- else -}}
                                HMI := "{{$os.HmiDb}}".o[{{$index}}]
                            {{- end}});
    END_REGION
{{end}}
END_FUNCTION`
)
