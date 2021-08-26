package environment

import "os"

type EnvVariable string

type Variable interface {
	Key() string
	Value() string
	Exists() bool
}

func (v EnvVariable) Key() string {
	return (string)(v)
}

func (v EnvVariable) Value() string {
	return os.Getenv(v.Key())
}

func (v EnvVariable) Exists() bool {
	return v.Value() != ""
}