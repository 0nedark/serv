package load

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

// Services defines the structure for array of services
type Services = []Service

// Groups defines structure of the groups section in serv yaml file
type Groups = map[string]Services

// ConfigFile defines structure of the serv yaml file
type ConfigFile struct {
	Order  []string `yaml:"order"`
	Groups Groups   `yaml:"groups"`
}
