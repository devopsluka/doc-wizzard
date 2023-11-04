package parser

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"errors"

	"github.com/devopsluka/doc-wizzard/docs"
	"github.com/devopsluka/doc-wizzard/types"
	"gopkg.in/yaml.v3"
)

func UnmarshalK8sObject(data []byte) (*types.K8sObject, error) {
	var rawObject map[string]interface{}
	if err := yaml.Unmarshal(data, &rawObject); err != nil {
		return nil, err
	}

	kind, ok := rawObject["kind"].(string)
	if !ok {
		return nil, errors.New("missing or invalid kind field")
	}

	var k8sObject types.K8sObject
	k8sObject.Kind = kind
	k8sObject.APIVersion, _ = rawObject["apiVersion"].(string)

	switch kind {
	case "Secret":
		fallthrough
	case "Service":
		var serviceSpec types.Service
		if err := yaml.Unmarshal(data, &serviceSpec); err != nil {
			return nil, err
		}
		k8sObject.Metadata = serviceSpec.Metadata
		k8sObject.Spec = &serviceSpec
	case "Deployment":
		var deployment types.Deployment
		if err := yaml.Unmarshal(data, &deployment); err != nil {
			return nil, err
		}
		k8sObject.Metadata = deployment.Metadata
		k8sObject.Spec = &deployment.Spec
	case "ConfigMap":
		var configMap types.ConfigMap
		if err := yaml.Unmarshal(data, &configMap); err != nil {
			return nil, err
		}
		k8sObject.Metadata = configMap.Metadata
		k8sObject.Spec = &configMap.Data
		//////////////////////////////////////////// PIPELINE ///////////////////////////////////////
	case "Pipeline":
		var pipeline types.Pipeline
		if err := yaml.Unmarshal(data, &pipeline); err != nil {
			return nil, err
		}
		k8sObject.Metadata = pipeline.Metadata
		k8sObject.Spec = &pipeline.PipelineSpec
	case "TriggerTemplate":
		var triggerTemplate types.TriggerTemplate
		if err := yaml.Unmarshal(data, &triggerTemplate); err != nil {
			return nil, err
		}
		k8sObject.Metadata = triggerTemplate.Metadata
		k8sObject.Spec = &triggerTemplate.Spec
	case "Route":
		var route types.Route
		if err := yaml.Unmarshal(data, &route); err != nil {
			return nil, err
		}
		k8sObject.Metadata = route.Metadata
		k8sObject.Spec = &route.RouteSpec
	case "EventListener":
		var eventListener types.EventListener
		if err := yaml.Unmarshal(data, &eventListener); err != nil {
			return nil, err
		}
		k8sObject.Metadata = eventListener.Metadata
		k8sObject.Spec = &eventListener.ElSpec
	default:
		return nil, errors.New("unsupported kind: " + kind)
	}

	return &k8sObject, nil
}

func ProcessK8sObject(obj *types.K8sObject, filePath string) error {
	doc, err := docs.GenerateK8sObjectDocumentation(*obj)
	if err != nil {
		return fmt.Errorf("failed to generate documentation for file %s: %w", filePath, err)
	}

	gpt, err := docs.CallGPT(*obj)
	if err != nil {
		return fmt.Errorf("failed to call OpenAI for documentation for file %s: %w", filePath, err)
	}

	fmt.Println(doc)
	fmt.Println(gpt)

	return nil
}

func ParseHelmChart(path string) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			obj, err := UnmarshalK8sObject(data)
			if err != nil {
				return fmt.Errorf("failed to unmarshal YAML in file %s: %w", filePath, err)
			}

			if err := ProcessK8sObject(obj, filePath); err != nil {
				return err
			}
		}

		return nil
	})
}
