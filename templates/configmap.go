package templates

const configMapTemplate = `
################## CONFIGMAP ##################
## ConfigMap: {{.Metadata.Name}}
- **Namespace**: {{.Metadata.Namespace}}
################## CONFIGMAP ##################
`
