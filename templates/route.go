package templates

const routeTemplate = `
################## ROUTE ##################
## Route: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}
- **Host**: {{.RouteSpec.Host}}
- **To**: {{ range $key, $value := .RouteSpec.ToService }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
- **Port**: {{ range $key, $value := .RouteSpec.RoutePort }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
- **Wildcard**: {{.RouteSpec.WildCard}}
- **Ingress**: {{ range $key, $value := .RouteStatus.RouteIngress }}
	- **{{ $key }}**: {{ $value }}
{{ end }}
################## ROUTE ##################
`
