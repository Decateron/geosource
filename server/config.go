package main

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Website Website
}

type Website struct {
	Url       string
	HttpPort  string
	HttpsPort string
}

// Returns a default configuration file
func NewConfig() *Config {
	return &Config{
		Website: Website{
			Url:       "localhost",
			HttpPort:  ":80",
			HttpsPort: ":443",
		},
	}
}

// Reads in config from a file with the given filename
func (config *Config) ReadFile(filename string) error {
	err := gcfg.ReadFileInto(config, filename)
	if err != nil {
		return err
	}
	return nil
}
