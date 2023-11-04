package types

type Deployment struct {
	Metadata   Metadata       `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
	Containers []Container    `yaml:"containers"`
}

type DeploymentSpec struct {
	Replicas int      `yaml:"replicas"`
	Template Template `yaml:"template"`
}

var _ Spec = (*DeploymentSpec)(nil)

type Spec2 struct {
	Containers []Container `yaml:"containers"`
}

type Template struct {
	Metadata Metadata          `yaml:"metadata"`
	Labels   map[string]string `yaml:"labels"`
	Spec     Spec2             `yaml:"spec"`
}

type Container struct {
	Name      string    `yaml:"name"`
	Image     string    `yaml:"image"`
	Ports     []Port    `yaml:"ports,omitempty"`
	Resources Resources `yaml:"resources,omitempty"`
}

type Resources struct {
	Limits   map[string]string `yaml:"limits"`
	Requests map[string]string `yaml:"requests"`
}

type Port struct {
	ContainerPort int `yaml:"containerPort"`
}
