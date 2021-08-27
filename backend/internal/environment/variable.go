package environment

import (
	"flag"
	"fmt"
	"os"
)

type EnvVariable string
type FlagVariable struct {
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
	vv := v.lookup()
	if vv == nil {
		return ""
	}
	return vv.Value.String()
}

func (v *FlagVariable) Value() interface{} {
	vv := &boolValue{}
	if err := vv.Set(v.String()); err != nil {
		return nil
	}
	return vv.value
}

type boolValue struct {
	value bool
}

func (val *boolValue) String() string {
	return fmt.Sprint(val.value)
}

func (val *boolValue) Set(value string) error {
	switch value {
	case "true":
		val.value = true
		return nil
	case "false":
		val.value = false
		return nil
	}
	return fmt.Errorf("could not parse %s to bool", value)
}
