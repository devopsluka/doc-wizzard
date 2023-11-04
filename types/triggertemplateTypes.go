package types

type TriggerTemplate struct {
	Metadata Metadata
	Spec     TriggerTemplateSpec
}

type TriggerTemplateSpec struct {
	Params      []TaskParam `yaml:"params"`
	PipelineRef PipelineRef `yaml:"pipelineRef"`
	Workspaces  []Workspace `yaml:"workspaces"`
}

type PipelineRef struct {
	Name string `yaml:"name"`
}

type Workspace struct {
	Name                string                `json:"name,omitempty"`
	ConfigMap           *ConfigMapWorkspace   `json:"configMap,omitempty"`
	Secret              *SecretWorkspace      `json:"secret,omitempty"`
	VolumeClaimTemplate *VolumeClaimWorkspace `json:"volumeClaimTemplate,omitempty"`
}

type ConfigMapWorkspace struct {
	Name string `yaml:"name, omitempty"`
}

type SecretWorkspace struct {
	Name string `yaml:"name, omitempty"`
}

type VolumeClaimWorkspace struct {
	Spec VolumeClaimTemplateSpec `json:"spec,omitempty"`
}

type VolumeClaimTemplateSpec struct {
	AccessModes []string `json:"accessModes,omitempty"`
	Resources   struct {
		Requests struct {
			Storage string `json:"storage,omitempty"`
		} `json:"requests,omitempty"`
	} `json:"resources,omitempty"`
	StorageClassName string `json:"storageClassName,omitempty"`
	VolumeMode       string `json:"volumeMode,omitempty"`
}

type Workspaces []Workspace
