package templates

const deploymentTemplate = `
################## DEPLOYMENT ##################
## Deployment: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}
- **Replicas**: {{ .DeploymentSpec.Replicas }}

### Containers ({{ len .DeploymentSpec.Template.Spec.Containers }}):
{{ range .DeploymentSpec.Template.Spec.Containers }}
- **Name**: {{ .Name }}
- **Image**: {{ .Image }} {{ if .Ports }}
- **Ports**:{{ range .Ports }}
	- **ContainerPort**: {{ .ContainerPort }} {{ end }}{{ end }}
{{ if .Resources }}
- **Resources**: {{ if .Resources.Limits }}
- **Limits**: {{ range $key, $value := .Resources.Limits }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
{{ end }}
{{ if .Resources.Requests }}
- **Requests**: {{ range $key, $value := .Resources.Requests }}
- **{{ $key }}**: {{ $value }}
{{ end }}
{{ end }}
{{ end }}
{{ end }}
################## DEPLOYMENT ##################
`
