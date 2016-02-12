package config

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Website  Website
	Google   Credentials
	Database Database
}

type Website struct {
	Url       string
	HttpPort  string
	HttpsPort string
}

type Credentials struct {
	ClientId     string
	ClientSecret string
	CallbackUrl  string
}

type Database struct {
	Host     string
	Database string
	User     string
	Password string
}

func New() *Config {
	return &Config{
		Website: Website{
			Url:       "localhost",
			HttpPort:  ":80",
			HttpsPort: ":443",
		},
		Google:   Credentials{},
		Database: Database{},
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