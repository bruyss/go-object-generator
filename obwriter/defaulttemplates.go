package obwriter

var Templates = map[string]string{
	"idbs":         idbTemplate,
	"measmon":      measmonTemplate,
	"digmon":       digmonTemplate,
	"valve":        valveTemplate,
	"controlValve": controlValveTemplate,
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

	measmonTemplate string = `FUNCTION {{.ObjectSettings.CallFc}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
VAR_INPUT
    {{.GeneralSettings.SecondPulse}}   : Bool;
    {{.GeneralSettings.Simulation}} : Bool;
END_VAR
{{$generalsettings := .GeneralSettings}}{{$objectsettings := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    REGION {{.InputMap.Tag}}: {{.InputMap.Description}}
        "{{.InputMap.IDB}}"(Tagname := '{{.InputMap.Tag}}',
                    Unit := '{{.InputMap.Unit}}',
                    SecPulse := {{$generalsettings.SecondPulse}},
                    Reset := TRUE,
                    Local := FALSE,
                    Simulation := {{$generalsettings.Simulation}},
                    AnalogInput := {{.InputMap.Input}},
                    LowLimit := {{.InputMap.LowLimit}},
                    HighLimit := {{.InputMap.HighLimit}},
                    {{if $generalsettings.Wincc}}HMI := "{{$objectsettings.HmiDb}}"."{{.InputMap.Tag}}"{{else}}HMI := "{{$objectsettings.HmiDb}}".o[{{$index}}]{{end}});
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
{{$generalsettings := .GeneralSettings}}{{$objectsettings := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    REGION {{.InputMap.Tag}}: {{.InputMap.Description}}
        "{{.InputMap.IDB}}"(Tagname := '{{.InputMap.Tag}}',
                    SecPulse := {{$generalsettings.SecondPulse}},
                    Reset := FALSE,
                    EnableAlarm := {{.InputMap.Alarm}},
                    InvertAlarm := {{.InputMap.InvertAlarm}},
                    Input := {{.InputMap.Input}},
                    Invert := {{.InputMap.Invert}},
                    HMI := "{{$objectsettings.HmiDb}}"."{{.InputMap.Tag}}");
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
{{$generalsettings := .GeneralSettings}}{{$objectsettings := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    REGION {{.InputMap.Tag}}: {{.InputMap.Description}}
        "{{.InputMap.IDB}}"(Tagname := '{{.InputMap.Tag}}',
                    SecPulse := {{$generalsettings.SecondPulse}},
                    Reset := FALSE,
                    Local := FALSE,
                    Simulation := {{$generalsettings.Simulation}},
                    Permit := TRUE,
                    NO := FALSE,
                    Activate := {{$generalsettings.TodoBit}},
                    FeedbackActivated := {{.InputMap.FBO}},
                    FeedbackDeactivated := {{.InputMap.FBC}},
                    MonitoringTimeAct := {{.InputMap.MonTimeOpen}},
                    MonitoringTimeDeact := {{.InputMap.MonTimeClose}}
                    Q_Activate := {{.InputMap.Output}},
                    {{if $generalsettings.Wincc}}HMI := "{{$objectsettings.HmiDb}}"."{{.InputMap.Tag}}"{{else}}HMI := "{{$objectsettings.HmiDb}}".o[{{$index}}]{{end}});
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
{{$generalsettings := .GeneralSettings}}{{$objectsettings := .ObjectSettings}}
BEGIN
{{range $index, $object := .Objects}}
    REGION {{.InputMap.Tag}}: {{.InputMap.Description}}
        "{{.InputMap.IDB}}"(Tagname := '{{.InputMap.Tag}}',
                    SecPulse := {{$generalsettings.SecondPulse}},
                    Reset := FALSE,
                    Local := FALSE,
                    Simulation := {{$generalsettings.Simulation}},
                    Permit := TRUE,
                    NoFeedback := {{.InputMap.NoFeedback}},
                    Feedback := {{.InputMap.Feedback}},
                    SP := {{$generalsettings.TodoReal}},
                    TimeMon := {{.InputMap.MonitoringTime}},
                    PQW := {{.InputMap.Output}},
                    {{if $generalsettings.Wincc}}HMI := "{{$objectsettings.HmiDb}}"."{{.InputMap.Tag}}"{{else}}HMI := "{{$objectsettings.HmiDb}}".o[{{$index}}]{{end}});
    END_REGION
{{end}}
END_FUNCTION`
)
