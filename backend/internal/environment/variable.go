package environment

import (
	"os"
)

type EnvVariable string

type Variable interface {
	Key() string
	Value() interface{}
	String() string
	Exists() bool
}

func (v EnvVariable) Key() string {
	return (string)(v)
}

func (v EnvVariable) Value() interface{} {
	return v.String()
}

func (v EnvVariable) String() string {
	return os.Getenv(v.Key())
}

func (v EnvVariable) Exists() bool {
	return v.String() != ""
}
