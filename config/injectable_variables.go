package config

import (
	"encoding/json"
	"fmt"
)

type InjectedVariables = map[string]string

func NewInjectedVariables(jsonBytes []byte) (*InjectedVariables, error) {
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
