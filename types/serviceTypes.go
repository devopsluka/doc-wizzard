package types

type Service struct {
	Metadata    Metadata    `yaml:"metadata"`
	ServiceSpec ServiceSpec `yaml:"spec"`
	Status      Status      `yaml:"status"`
}

type ServiceSpec struct {
	Metadata Metadata          `yaml:"metadata"`
	Selector map[string]string `yaml:"selector"`
	Ports    []ServicePort     `yaml:"ports"`
	Type     string            `yaml:"type"`
}

type ServicePort struct {
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
	Name       string `yaml:"name"`
}

type Status struct {
	LoadBalancer LoadBalancer `yaml:"loadBalancer"`
}

type LoadBalancer struct {
	Ingress []Ingress `yaml:"ingress"`
}

type Ingress struct {
	IP string `yaml:"ip"`
}
