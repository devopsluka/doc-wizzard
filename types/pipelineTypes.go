package types

type Pipeline struct {
	Metadata     Metadata     `yaml:"metadata"`
	PipelineSpec PipelineSpec `yaml:"spec"`
}

type PipelineSpec struct {
	PipelineParams []PipelineParam `yaml:"params"`
	Tasks          []Task          `yaml:"tasks"`
	Workspaces     []Workspace     `yaml:"workspaces"`
}

type PipelineParam struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Default     string `yaml:"default"`
	Description string `yaml:"description"`
}

type Task struct {
	Name       string      `yaml:"name"`
	TaskParams []TaskParam `yaml:"params"`
	Steps      []Step      `yaml:"steps"` // TASK ONLY
	RunAfter   []string    `yaml:"runAfter"`
	TaskRef    TaskRef     `yaml:"taskRef"`
	Workspaces []Workspace `yaml:"workspaces"`
}

type TaskParam struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type TaskRef struct {
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
}

// TASK ONLY
type Step struct {
	Name   string            `yaml:"name"`
	Script string            `yaml:"script"`
	Image  string            `yaml:"image"`
	Env    map[string]string `yaml:"env"`
}
