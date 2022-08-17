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

	measmonTemplate string = `FUNCTION {{.FcName}} : Void
{ S7_Optimized_Access := 'TRUE'}
VERSION : 0.1
    VAR_INPUT
        iSecPulse   : Bool;
        iSimulation : Bool;
    END_VAR

BEGIN
{{range .Objects}}
    REGION {{.InputMap.Tag}}: {{.InputMap.Description}}
        "{{.InputMap.IDB}}" := '{{.InputMap.Tag}}',
                    Unit := '{{.InputMap.Unit}}',
                    SecPulse := iSecPulse,
                    Reset := TRUE,
                    Local := FALSE,
                    Simulation := iSimulation,
                    AnalogInput := {{.InputMap.Input}},
                    LowLimit := {{.InputMap.LowLimit}},
                    HighLimit := {{.InputMap.HighLimit}},
                    HMI := "HMI_Measmon"."{{.InputMap.Tag}}");
    END_REGION
{{end}}
END_FUNCTION`
)
