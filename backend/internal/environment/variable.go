package environment

import (
	"os"
)

// EnvVariable is a key of an environment variable
type EnvVariable string

// Variable is a key/value association of an application configuration variable
type Variable interface {
	Key() string
	Value() interface{}
	String() string
	Exists() bool
}

// Key provides the name of the EnvVariable
func (v EnvVariable) Key() string {
	return (string)(v)
}

// Value provides the value of the EnvVariable
func (v EnvVariable) Value() interface{} {
	return v.String()
}

// String provides the value via lookup from the underlying operating system environment variables
func (v EnvVariable) String() string {
	return os.Getenv(v.Key())
}

func (v EnvVariable) Exists() bool {
	return v.String() != ""
}
