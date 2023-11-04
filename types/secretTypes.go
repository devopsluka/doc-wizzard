package types

type Secret struct {
	Metadata Metadata `yaml:"metadata"`
	Data     Data     `yaml:"data"`
}
