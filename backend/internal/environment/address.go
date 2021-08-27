package environment

import "fmt"

func (env *Environment) HttpAddress() string {
	return fmt.Sprintf(":%s", env.PortHttp.String())
}

func (env *Environment) TcpAddress() string {
	return fmt.Sprintf(":%s", env.PortTcp.String())
}
