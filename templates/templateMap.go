package templates

import (
	"html/template"
)

var TemplatesMap = map[string]*template.Template{
	"Deployment":      template.Must(template.New("deployment").Parse(deploymentTemplate)),
	"Service":         template.Must(template.New("service").Parse(serviceTemplate)),
	"ConfigMap":       template.Must(template.New("configMap").Parse(configMapTemplate)),
	"Pipeline":        template.Must(template.New("pipeline").Parse(pipelineTemplate)),
	"TriggerTemplate": template.Must(template.New("triggerTemplate").Parse(triggerTemplate)),
	"Route":           template.Must(template.New("route").Parse(routeTemplate)),
	"EventListener":   template.Must(template.New("eventListener").Parse(eventListenerTemplate)),
}
