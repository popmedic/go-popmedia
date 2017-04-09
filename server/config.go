package server

import (
	"encoding/json"
	"io/ioutil"
)

var MainConfig *Config

type Config struct {
	Port string
	Root string
	MediaExt []string
	DirectoryImage string
}

func newConfig(path string) (*Config, error) {
	jsonData, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}

	cfg := &Config{
		Port: "8080",
		Root: "/",
		MediaExt: []string{},
	}

	err = json.Unmarshal(jsonData, cfg)
	if nil != err {
		return nil, err
	}
	return cfg, nil
}

func init(){
	var err error
	MainConfig, err = newConfig("config.json")
	if nil != err {
		panic("unable to open config.json")
	}
}