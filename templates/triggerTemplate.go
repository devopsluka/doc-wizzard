package templates

const triggerTemplate = `
################## TRIGGER TEMPLATE ##################
## TriggerTemplate: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}
- **Params**: {{ range $key, $value := .TriggerTemplateSpec.Params }}
- **{{ $key }}**: {{ $value }}
{{ end }}
- **PipelineRef**: {{ range $key, $value := .TriggerTemplateSpec.PipelineRef }}
- **{{ $key }}**: {{ $value }}
{{ end }}
- **Workspaces**: {{ range $key, $value := .TriggerTemplateSpec.Workspaces }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
################## TRIGGER TEMPLATE ##################
`
