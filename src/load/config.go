package load

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var read = ioutil.ReadFile

// Config the specified yml file
func Config(file string) (ConfigFile, error) {
	raw, err := read(file)
	if err != nil {
		return ConfigFile{}, err
	}

	config := ConfigFile{}
	if err = yaml.Unmarshal(raw, &config); err != nil {
		return ConfigFile{}, err
	}

	return config, nil
}
