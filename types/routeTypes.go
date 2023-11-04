package types

type Route struct {
	Metadata    Metadata    `yaml:"metadata"`
	RouteSpec   RouteSpec   `yaml:"spec"`
	RouteStatus RouteStatus `yaml:"status"`
}

type RouteSpec struct {
	Host      string    `yaml:"host"`
	ToService ToService `yaml:"to"`
	RoutePort RoutePort `yaml:"port"`
	WildCard  string    `yaml:"wildcardPolicy"`
}

type ToService struct {
	Kind   string `yaml:"kind"`
	Name   string `yaml:"name"`
	Weight int    `yaml:"weight"`
}

type RoutePort struct {
	TargetPort string `yaml:"targetPort"`
}

type RouteStatus struct {
	RouteIngress []RouteIngress `yaml:"ingress"`
}

type RouteIngress struct {
	Host           string `yaml:"host"`
	RouterName     string `yaml:"routerName"`
	WildcardPolicy string `yaml:"wildcardPolicy"`
	RouterHostName string `yaml:"routerCanonicalHostname"`
}
