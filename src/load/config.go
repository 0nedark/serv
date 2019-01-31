package load

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Command defines structure of bash command
type Command struct {
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

// Healthcheck defines structure of health check command
type Healthcheck struct {
	Command `yaml:",inline"`
	Timeout int `yaml:"timeout"`
	Sleep   int `yaml:"sleep"`
}

// Repository defines structure of single repository
type Repository struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

// Service defines structure of a single project
type Service struct {
	Repository     `yaml:",inline"`
	Command          Command       `yaml:"command"`
	Healthchecks   []Healthcheck `yaml:"healthchecks"`
	Postconditions []Command     `yaml:"postconditions"`
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
