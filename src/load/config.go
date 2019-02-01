package load

import (
	"io/ioutil"

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

// GetConfig the specified yml file into the provided structure
func GetConfig(file string, input *Config) error {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(raw, input); err != nil {
		return err
	}

	return nil
}
