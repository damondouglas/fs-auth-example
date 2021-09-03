package environment

import "fmt"

// HttpAddress provides the address to bind an HTTP service from the Environment
func (env *Environment) HttpAddress() string {
	return fmt.Sprintf(":%s", env.PortHttp.String())
}

// TcpAddress provides the address to bind an TCP service from the Environment
func (env *Environment) TcpAddress() string {
	return fmt.Sprintf(":%s", env.PortTcp.String())
}
