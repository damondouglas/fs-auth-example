package environment

import (
	"encoding/json"
	"fs-auth-example/backend/internal/gcloud"
)

var (
	DefaultEnvironment = &Environment{
		ClientId: EnvClientId,
		PortHttp: EnvHttpPort,
		PortTcp:  EnvTcpPort,
		project: EnvProject,
	}
)

type Environment struct {
	ClientId Variable
	PortHttp Variable
	PortTcp  Variable
	project Variable
}

func (env *Environment) String() string {
	s := map[string]string{}
	for _, v := range []Variable{
		env.ClientId,
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

func (env *Environment) Project() (string, error) {
	if env.project.Exists() {
		return env.project.String(), nil
	}
	return gcloud.Config.Project()
}
