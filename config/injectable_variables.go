package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InjectedVariables = map[string]string

func ReadInjectedVariablesFromFile(filePath string) (*InjectedVariables, error) {
	varFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	varsFromFile, err := injectedVariablesFromJsonBytes(varFile)
	if err != nil {
		return nil, err
	}

	return varsFromFile, nil
}

func injectedVariablesFromJsonBytes(jsonBytes []byte) (*InjectedVariables, error) {
	var userInjectedVars = &map[string]string{}

	ok := json.Valid(jsonBytes)
	if !ok {
		return nil, fmt.Errorf("provided configuration is not a valid json")
	}

	if err := json.Unmarshal(jsonBytes, userInjectedVars); err != nil {
		return nil, err
	}

	return userInjectedVars, nil
}
