package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type config struct {
	ListenAddress string `json:"listen_address"`
	Backends      map[string]struct {
		Destination string `json:"destination"`
		Certificate string `json:"certificate"`
		Key         string `json:"key"`
	}
}

var Config config

func (c *config) Load() {
	data, _ := ioutil.ReadFile("config.json")
	_ = json.Unmarshal(data, c)
}

func (c *config) GetCertificatesForDomain(domain string) (string, string, error) {
	if value, ok := Config.Backends[domain]; ok {
		return value.Certificate, value.Key, nil
	}
	return "", "", errors.New("config not found: " + domain)
}
func (c *config) GetBackendForDomain(domain string) (string, error) {
	if value, ok := Config.Backends[domain]; ok {
		return value.Destination, nil
	}
	return "", errors.New("config not found: " + domain)
}
