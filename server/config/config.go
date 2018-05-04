package config

import (
	"encoding/json"
	"io/ioutil"
)

var MainConfig = &Config{}

type Config struct {
	Port           string
	Root           string
	Host           string
	MediaExt       []string
	DirectoryImage string
	FileImage      string
	LogFile        string
	UseHTMLLogger  bool
	HTMLLoggerPort string
	path           string
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{
		Port:           "8080",
		Root:           "/",
		MediaExt:       []string{},
		UseHTMLLogger:  false,
		HTMLLoggerPort: "6060",
		path:           path,
	}

	err := cfg.LoadConfig(path)
	if nil != err {
		return nil, err
	}
	return cfg, nil
}

func (cfg *Config) LoadConfig(path string) error {
	cfg.path = path
	return cfg.ReloadConfig()
}

func (cfg *Config) ReloadConfig() error {
	jsonData, err := ioutil.ReadFile(cfg.path)
	if nil != err {
		return err
	}

	return json.Unmarshal(jsonData, cfg)
}
