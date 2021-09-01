package testdata

import (
	"bytes"
	"github.com/joho/godotenv"
	"os"
)

func LoadDotEnv() error {
	kv, err := godotenv.Parse(bytes.NewReader(dotEnv))
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
