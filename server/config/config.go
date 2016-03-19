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
	// Url from which the site can be accessed.
	Url string
	// Port over which HTTP is served
	HttpPort string
	// Port over which HTTPS is served
	HttpsPort string
	// Path to the TLS certificate
	Cert string
	// Path to the TLS key
	Key string
}

type Credentials struct {
	ClientID     string
	ClientSecret string
	CallbackUrl  string
}

type Database struct {
	Host     string
	Database string
	User     string
	Password string
}

// Returns a new config file with default values set where applicable.
func New() *Config {
	return &Config{
		Website: Website{
			Url:       "localhost",
			HttpPort:  ":80",
			HttpsPort: ":443",
			Cert:      "cert.pem",
			Key:       "key.pem",
		},
	}
}

// Reads in config information from given file assumed to be in gcfg format,
// overwriting any existing values. Returns an error if the file cannot be read,
// nil otherwise.
func (config *Config) ReadFile(filename string) error {
	err := gcfg.ReadFileInto(config, filename)
	if err != nil {
		return err
	}
	return nil
}
