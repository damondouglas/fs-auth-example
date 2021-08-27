package environment

import (
	"flag"
	"fmt"
	"os"
)

type EnvVariable string
type FlagVariable struct {
	value *bool
	*flag.Flag
}

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

func (v *FlagVariable) Key() string {
	return v.Name
}

func (v *FlagVariable) lookup() *flag.Flag {
	if !flag.Parsed() {
		return nil
	}
	return flag.Lookup(v.Key())
}

func (v *FlagVariable) Exists() bool {
	return v.lookup() != nil
}

func (v *FlagVariable) String() string {
	return fmt.Sprint(*v.value)
}

func (v *FlagVariable) Value() interface{} {
	return v.value
}
