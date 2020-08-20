package main

import (
	"errors"
)

// Config is web quic server structure config.
type Config struct {
	Address string `json:"address"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`

	APIKey string `json:"api_key"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Address == rhs.Address &&
		c.Cert == rhs.Cert &&
		c.Key == rhs.Key)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	if c.Address, ok = cAddress.(string); !ok {
		return errors.New("key address invalid. must be string")
	}

	cCert, ok := fconf["cert"]
	if !ok {
		return errors.New("missing key cert")
	}
	if c.Cert, ok = cCert.(string); !ok {
		return errors.New("key cert invalid. must be string")
	}

	cKey, ok := fconf["key"]
	if !ok {
		return errors.New("missing key key")
	}
	if c.Key, ok = cKey.(string); !ok {
		return errors.New("key key invalid. must be string")
	}

	cAPIKey, ok := fconf["api_key"]
	if !ok {
		return errors.New("missing key api_key")
	}
	if c.APIKey, ok = cAPIKey.(string); !ok {
		return errors.New("key api_key invalid. must be string")
	}

	return nil
}
