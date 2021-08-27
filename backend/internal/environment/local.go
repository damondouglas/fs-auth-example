package environment

import (
	"bytes"
	"fs-auth-example/backend/testdata"
	"os"

	"github.com/joho/godotenv"
)

func (env *Environment) IsLocal() bool {
	v := env.Local.Value()
	if v == nil {
		return false
	}
	if vv, ok := v.(bool); ok {
		return vv
	}
	return false
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
