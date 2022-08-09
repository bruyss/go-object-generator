package obwriter

var Templates = map[string]string{
	"callfc":  callFcTemplate,
	"measmon": measmonTemplate,
}

const (
	callFcTemplate string = `{{range .}}
    {{.Tag}}
{{end}}`
)

const (
	measmonTemplate string = `{{range .}}
REGION {{.Tag}}: {{.Description}}

        "IDB" := '{{.Tag}}',
                       Unit := '{{.Unit}}',
                       SecPulse := iSecPulse,
                       Reset := TRUE,
                       Local := FALSE,
                       Simulation := iSimulation,
                       AnalogInput := {{.Tag}},
                       LowLimit := {{.LowLimit}},
                       HighLimit := {{.HighLimit}},
                       HMI := "HMI_Measmon"."{{.Tag}}");
{{end}}`
)
