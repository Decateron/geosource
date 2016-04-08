// Package config provides a common structure that is shared throughout the
// program to enable alterable settings.
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
	// URL from which the site can be accessed.
	URL string
	// Port over which HTTP is served
	HTTPPort string
	// Port over which HTTPS is served
	HTTPSPort string
	// Path to the TLS certificate
	Cert string
	// Path to the TLS key
	Key string
}

// Credentials is a generalized OAuth token structure which is shared by various
// login providers such as Google, Facebook, and Twitter.
type Credentials struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// Database contains information related to the location of the database.
type Database struct {
	Host     string
	Database string
	User     string
	Password string
}

// New constructs a new config file with default values set where applicable.
func New() *Config {
	return &Config{
		Website: Website{
			URL:       "localhost",
			HTTPPort:  ":80",
			HTTPSPort: ":443",
			Cert:      "cert.pem",
			Key:       "key.pem",
		},
	}
}

// ReadFile reads in config information from the file with the given name,
// assumed to be in gcfg format, overwriting any existing values. Returns an
// error if the file cannot be read, nil otherwise.
func (config *Config) ReadFile(filename string) error {
	err := gcfg.ReadFileInto(config, filename)
	if err != nil {
		return err
	}
	return nil
}
