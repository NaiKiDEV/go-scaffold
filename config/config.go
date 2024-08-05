package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	Name         string `json:"name"`
	Template     string `json:"template"`
	TemplateFile string `json:"templateFile"`
}

type Directory struct {
	Name           string      `json:"name"`
	Files          []File      `json:"files"`
	SubDirectories []Directory `json:"subDirectories"`
}

type Config struct {
	TemplateVarsPath string      `json:"templateVarsPath"`
	TemplateRootPath string      `json:"templateRootPath"`
	DirectoryConfig  []Directory `json:"directoryConfig"`
}

func ReadConfigFromFile(filePath string) (*Config, error) {
	cfgFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cfgFromFile, err := configFromJsonBytes(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfgFromFile, nil

}

func configFromJsonBytes(jsonBytes []byte) (*Config, error) {
	var cfg = &Config{}

	ok := json.Valid(jsonBytes)
	if !ok {
		return nil, fmt.Errorf("provided file is not a valid json")
	}

	if err := json.Unmarshal(jsonBytes, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
