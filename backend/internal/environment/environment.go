package environment

import (
	"encoding/json"
	"fmt"
	"fs-auth-example/backend/internal/gcloud"
	"strings"
)

var (
	DefaultEnvironment = &Environment{
		ClientId: EnvClientId,
		PortHttp: EnvHttpPort,
		PortTcp:  EnvTcpPort,
		Verbose:  FlagVerbose,
		Local:    FlagLocal,
	}
)

type Environment struct {
	ClientId Variable
	PortHttp Variable
	PortTcp  Variable
	Verbose  Variable
	Local    Variable
}

func (env *Environment) String() string {
	s := map[string]interface{}{}
	for _, v := range []Variable{
		env.ClientId,
		env.Local,
		env.PortHttp,
		env.PortTcp,
		env.Verbose,
	} {
		if v == nil {
			continue
		}
		switch v.(type) {
		case EnvVariable:
			s[v.Key()] = v.String()
		case *FlagVariable:
			s[v.Key()] = v.Value()
		}
	}
	b, _ := json.Marshal(s)
	return string(b)
}

func Missing(vars ...Variable) error {
	var missingEnv []string
	var missingFlag []string
	for _, k := range vars {
		if !k.Exists() {
			switch k.(type) {
			case EnvVariable:
				missingEnv = append(missingEnv, k.Key())
			case *FlagVariable:
				key := fmt.Sprintf("-%s", k.Key())
				missingFlag = append(missingFlag, key)
			}
		}
	}

	if len(missingEnv)+len(missingFlag) == 0 {
		return nil
	}

	missingEnvMessage := "none"
	missingFlagMessage := "none"

	if len(missingEnv) > 0 {
		missingEnvMessage = strings.Join(missingEnv, " | ")
	}

	if len(missingFlag) > 0 {
		missingFlagMessage = strings.Join(missingFlag, " | ")
	}

	return fmt.Errorf("fatal: required environment or flag variables are missing (note: requires flag.Parse()); environment variables: %s; flag variables: %s", missingEnvMessage, missingFlagMessage)
}

func (env *Environment) Project() (string, error) {
	return gcloud.Config.Project()
}

func (env *Environment) IsVerbose() bool {
	v := env.Verbose.Value()
	if v == nil {
		return false
	}
	if vv, ok := v.(bool); ok {
		return vv
	}
	return false
}
