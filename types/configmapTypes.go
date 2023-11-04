package types

type ConfigMap struct {
	Metadata Metadata `yaml:"metadata"`
	Data     Data     `yaml:"data"`
}

type Data struct {
	Keys map[string]interface{} `yaml:"data"`
}
