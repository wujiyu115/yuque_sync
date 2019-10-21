package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
)

//SyncConfig of sync
type SyncConfig struct {
	Token       string `json:"token" yaml:"token"`
	Login       string `json:"login" yaml:"login"`
	Repo        string `json:"repo" yaml:"repo"`
	PostPath    string `json:"postPath" yaml:"postPath"`
	CachePath   string `json:"cachePath" yaml:"cachePath"`
	MdFormat    string `json:"mdFormat" yaml:"mdFormat"`
	Concurrency int    `json:"concurrency" yaml:"concurrency"`
	OnlyPub     bool   `json:"onlyPub" yaml:"onlyPub"`
}

var defaultConfig = SyncConfig{
	Token:       "",
	Login:       "",
	Repo:        "",
	PostPath:    "source/_posts/yuque",
	CachePath:   "yuque.json",
	MdFormat:    "title",
	Concurrency: 5,
	OnlyPub:     false,
}

func loadJSON(cfg *SyncConfig) error {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return json.Unmarshal(byteValue, &cfg)
}

func loadYaml(cfg *SyncConfig) error {
	jsonFile, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return yaml.Unmarshal(byteValue, &cfg)
}

//LoadConfig of sync
func LoadConfig() SyncConfig {
	var (
		cfg SyncConfig
	)
	if err := loadJSON(&cfg); err != nil {
		loadYaml(&cfg)
	}
	copier.Copy(&defaultConfig, &cfg)
	return defaultConfig
}
