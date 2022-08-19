package obwriter

var Templates = map[string]string{
	"idbs":    idbTemplate,
	"measmon": measmonTemplate,
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
            "{{.InputMap.IDB}}" := '{{.InputMap.Tag}}',
                        Unit := '{{.InputMap.Unit}}',
                        SecPulse := {{$generalsettings.SecondPulse}},
                        Reset := TRUE,
                        Local := FALSE,
                        Simulation := {{$generalsettings.Simulation}},
                        AnalogInput := {{.InputMap.Input}},
                        LowLimit := {{.InputMap.LowLimit}},
                        HighLimit := {{.InputMap.HighLimit}},
                        HMI := "{{$objectsettings.HmiDb}}"."{{.InputMap.Tag}}");
        END_REGION
    {{end}}
    END_FUNCTION`
)
