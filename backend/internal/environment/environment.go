package environment

import (
	"fmt"
	"strings"
)

type Environment struct {
	Port Variable
}

func New(requiredVars ...Variable) (*Environment, error) {
	missingVars := missing(requiredVars...)
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("fatal: environment variables required but missing: %s", strings.Join(missingVars, " | "))
	}
	return &Environment{
		Port: EnvPort,
	}, nil
}

func missing(vars ...Variable) []string {
	var result []string
	for _, k := range vars {
		if !k.Exists() {
			result = append(result, k.Key())
		}
	}
	return result
}
