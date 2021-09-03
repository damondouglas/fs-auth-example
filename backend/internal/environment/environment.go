package environment

import (
	"encoding/json"
	"fs-auth-example/backend/internal/gcloud"
)

var (
	// DefaultEnvironment is a convenience wrapper instance of an Environment
	DefaultEnvironment = &Environment{
		PortHttp: EnvHttpPort,
		PortTcp:  EnvTcpPort,
		project: EnvProject,
	}
)

// Environment holds common environment variables for the service
type Environment struct {
	PortHttp Variable
	PortTcp  Variable
	project Variable
}

// String returns a string representation of an Environment
func (env *Environment) String() string {
	s := map[string]string{}
	for _, v := range []Variable{
		env.PortHttp,
		env.PortTcp,
	} {
		if v == nil {
			continue
		}
		s[v.Key()] = v.String()
	}
	b, _ := json.Marshal(s)
	return string(b)
}

// Project returns an environment variable value associated with EnvProject
// or attempts to acquire locally from a gcloud installation
func (env *Environment) Project() (string, error) {
	if env.project.Exists() {
		return env.project.String(), nil
	}
	return gcloud.Config.Project()
}
