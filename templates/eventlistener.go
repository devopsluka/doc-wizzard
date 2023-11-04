package templates

const eventListenerTemplate = `
################## EVENT LISTENER ##################
## EventListener: {{ .Metadata.Name }}
- **Namespace**: {{ .Metadata.Namespace }}

### Spec:
- **Service Account Name**: {{ .ElSpec.ServiceAccountName }}

### Triggers ({{ len .ElSpec.Triggers }}):
{{ range .ElSpec.Triggers }}
- **Bindings**:
  {{- range .Bindings }}
    - **Kind**: {{ .Kind }}
    - **Ref**: {{ .Ref }}
  {{- end }}
- **Interceptors**:
  {{- range .Interceptors }}
    - **Params**:
      {{- range .Params }}
        - **Name**: {{ .Name }}
        - **Value**: 
          {{- if .Value }}
          {{- range $key, $value := .Value }}
            - **{{ $key }}**: {{ $value }}
          {{- end }}
          {{- end }}
      {{- end }}
    - **Ref**:
      - **Kind**: {{ .InterceptorRef.Kind }}
      - **Name**: {{ .InterceptorRef.Name }}
  {{- end }}
- **Template Ref**:
    - **Ref**: {{ .TemplateRef.Ref }}
{{ end }}

### Status:
- **Address**:
    - **URL**: {{ .ElStatus.Address.URL }}
- **Configuration**:
    - **Generated Name**: {{ .ElStatus.Configuration.GeneratedName }}
################## EVENT LISTENER ##################
`
