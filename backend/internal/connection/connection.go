package connection

import (
	"fs-auth-example/backend/internal/environment"
)

type Connection struct {
	env *environment.Environment
}

func FromEnvironment(env *environment.Environment) *Connection {
	return &Connection{
		env: env,
	}
}
