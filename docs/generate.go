package docs

import (
	"bytes"
	"fmt"

	"github.com/devopsluka/doc-wizzard/templates"
	"github.com/devopsluka/doc-wizzard/types"
)

func GenerateK8sObjectDocumentation(obj types.K8sObject) (string, error) {
	tmpl, ok := templates.TemplatesMap[obj.Kind]
	if !ok {
		return "", fmt.Errorf("unsupported Kubernetes object kind: %s", obj.Kind)
	}

	var doc bytes.Buffer
	switch obj.Kind {
	case "Deployment":
		if spec, ok := obj.Spec.(*types.DeploymentSpec); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata       types.Metadata
				DeploymentSpec types.DeploymentSpec
			}{obj.Metadata, *spec}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for Deployment: %T", obj.Spec)
		}
	case "Service":
		if service, ok := obj.Spec.(*types.Service); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata types.Metadata
				Spec     types.ServiceSpec
				Status   types.Status
			}{obj.Metadata, service.ServiceSpec, service.Status}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for Service: %T", obj.Spec)
		}
	case "ConfigMap":
		if data, ok := obj.Spec.(*types.Data); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata types.Metadata
				Data     types.Data
			}{obj.Metadata, *data}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for ConfigMap: %T", obj.Spec)
		}
		//////////////////////////////////// PIPELINE ///////////////////////////////////////
	case "Pipeline":
		if pipeline, ok := obj.Spec.(*types.PipelineSpec); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata     types.Metadata
				PipelineSpec types.PipelineSpec
			}{obj.Metadata, *pipeline}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for Pipeline: %T", obj.Spec)
		}
	case "TriggerTemplate":
		if triggerTemplate, ok := obj.Spec.(*types.TriggerTemplateSpec); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata            types.Metadata
				TriggerTemplateSpec types.TriggerTemplateSpec
			}{obj.Metadata, *triggerTemplate}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for TriggerTemplate: %T", obj.Spec)
		}
	case "Route":
		if route, ok := obj.Spec.(*types.RouteSpec); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata  types.Metadata
				RouteSpec types.RouteSpec
			}{obj.Metadata, *route}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for Route: %T", obj.Spec)
		}
	case "EventListener":
		if eventListener, ok := obj.Spec.(*types.ElSpec); ok {
			if err := tmpl.Execute(&doc, struct {
				Metadata           types.Metadata
				EventListenerSpecs types.ElSpec
			}{obj.Metadata, *eventListener}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("invalid spec type for EventListener: %T", obj.Spec)
		}
	default:
		return "", fmt.Errorf("unsupported Kubernetes object kind: %s", obj.Kind)

	}

	Append2File("output/documentation.txt", doc.String())

	return doc.String(), nil
}
