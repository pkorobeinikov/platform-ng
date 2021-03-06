package spec

import (
	"os"

	"gopkg.in/yaml.v3"
)

const File = "platform.yaml"

type (
	Spec struct {
		Name        string                       `yaml:"name"`
		Namespace   string                       `yaml:"namespace"`
		Environment map[string]map[string]string `yaml:"environment"`
		Component   []*Component                 `yaml:"component"`
	}

	Component struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}
)

func Load(file string) (*Spec, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(b, &spec)
	if err != nil {
		return nil, err
	}

	return &spec, nil
}
