package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonConf map[string]interface{}

// Config struct have all the mailing templates info.
type Config struct {
	filePath string
}

// NewConfig reads and loads the configuation of all the inputs and outputs.
func NewConfig(filePath string) *Config {
	return &Config{
		filePath: filePath,
	}
}

// Load loads the configuration file.
func (c *Config) Load(key string) (interface{}, error) {
	content, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return nil, err
	}

	var object jsonConf
	err = json.Unmarshal(content, &object)
	if err != nil {
		return nil, err
	}

	val, ok := object[key]
	if !ok {
		return nil, fmt.Errorf("Invalid %s key in configuration", key)
	}

	return val, nil
}
