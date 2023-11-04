package templates

const pipelineTemplate = `
################## PIPELINE ##################
## Pipeline: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}

### Spec:
- **Params**: {{ range $key, $value := .Spec.Params }}
- **{{ $key }}**: {{ $value }}
{{ end }}
- **PipelineSpec**: {{ range $key, $value := .Spec.PipelineSpec }}
- **{{ $key }}**: {{ $value }}
{{ end }}

### Tasks ({{ len .Spec.Tasks }}):
{{ range .Spec.Tasks }}
- **Name**: {{ .Name }}
- **Params**: {{ range $key, $value := .TaskParams }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
- **RunAfter**: {{ range $key, $value := .RunAfter }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
- **TaskRef**: {{ range $key, $value := .TaskRef }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
- **Workspaces**: {{ range $key, $value := .Workspaces }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
{{ end }}
################## PIPELINE ##################
`
