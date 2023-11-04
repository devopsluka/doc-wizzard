package templates

const serviceTemplate = `
################## SERVICE ##################
## Service: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}
- **Type**: {{ .Spec.Type}}
### Selector: {{ range $key, $value := .Spec.Selector }}
- **{{ $key }}**: {{ $value }}
{{ end }}
### Ports ({{ len .Spec.Ports }}):
{{ range .Spec.Ports }}
- **Name**: {{ .Name }}
- **Port**: {{ .Port }}
- **TargetPort**: {{ .TargetPort }}
{{ end }}
### Status:
{{ if .Status.LoadBalancer.Ingress }}loadBalancer:{{ range .Status.LoadBalancer.Ingress }}
- **IP**: {{ .IP }}
{{ end }}
{{ else }}
- **IP**: Pending
{{ end }}
################## SERVICE ##################
`
