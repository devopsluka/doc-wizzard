package types

type EventListener struct {
	Metadata Metadata `yaml:"metadata"`
	ElSpec   ElSpec   `yaml:"spec"`
	ElStatus ElStatus `yaml:"status"`
}

type ElSpec struct {
	ServiceAccountName string    `yaml:"serviceAccountName"`
	Triggers           []Trigger `yaml:"triggers"`
}

type Trigger struct {
	Bindings     []Binding      `yaml:"bindings"`
	Interceptors []*Interceptor `yaml:"interceptors"`
	TemplateRef  TemplateRef    `yaml:"template"`
}

type Binding struct {
	Kind string `yaml:"kind"`
	Ref  string `yaml:"ref"`
}

type Interceptor struct {
	Params []*InterceptorParam `yaml:"params"`
}

type InterceptorParam struct {
	Name           string         `yaml:"name"`
	Value          string         `yaml:"value"`
	InterceptorRef InterceptorRef `yaml:"ref"`
}

type InterceptorRef struct {
	Kind string `yaml:"kind"`
	Name string `yaml:"name"`
}

type TemplateRef struct {
	Ref string `yaml:"ref"`
}

type ElStatus struct {
	Address       Address       `yaml:"address"`
	Configuration Configuration `yaml:"configuration"`
}

type Address struct {
	URL string `yaml:"url"`
}

type Configuration struct {
	GeneratedName string `yaml:"generatedName"`
}
