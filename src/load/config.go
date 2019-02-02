package load

import (
	yaml "gopkg.in/yaml.v2"
)

// Postcondition defines structure of post condition command
type Postcondition struct {
	Command string `yaml:"command"`
}

// Healthcheck defines structure of health check command
type Healthcheck struct {
	Command string `yaml:"command"`
	Timeout int    `yaml:"timeout"`
	Sleep   int    `yaml:"sleep"`
}

// Repository defines structure of single repository
type Repository struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

// Service defines structure of a single project
type Service struct {
	Repository     `yaml:",inline"`
	Command        string          `yaml:"command"`
	Healthchecks   []Healthcheck   `yaml:"healthchecks"`
	Postconditions []Postcondition `yaml:"postconditions"`
}

// Groups defines structure of the groups section in serv yaml file
type Groups = map[string][]Service

// Config defines structure of the serv yaml file
type Config struct {
	Order  []string `yaml:"order"`
	Groups Groups   `yaml:"groups"`
}

// ReadFileFunc defines read file function signature
type ReadFileFunc = func(string) ([]byte, error)

// ConfigFunc defines the config function signature
type ConfigFunc = func(string, ReadFileFunc) (Config, error)

// NewConfig the specified yml file
func NewConfig(file string, readFile ReadFileFunc) (Config, error) {
	raw, err := readFile(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if err = yaml.Unmarshal(raw, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
