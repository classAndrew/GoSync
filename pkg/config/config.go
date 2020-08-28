package config

import (
	"io/ioutil"
)

// Config struct for configuration
type Config struct {
	RootDir string
	Port    int
}

//ReadConfig reads configuration
func ReadConfig() (Config, error) {
	res, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return "", err
	}
	return Config{}, nil
}
