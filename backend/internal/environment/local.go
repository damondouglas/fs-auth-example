package environment

import (
	"bytes"
	"fs-auth-example/backend/testdata"
	"os"

	"github.com/joho/godotenv"
)

func (env *Environment) IsLocal() bool {
	local := env.Local.(*FlagVariable)
	return *local.value
}

func LoadLocalEnv() error {
	kv, err := godotenv.Parse(bytes.NewBuffer(testdata.DotEnv))
	if err != nil {
		return err
	}
	for k, v := range kv {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}
