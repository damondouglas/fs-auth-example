package environment

var (
	// EnvHttpPort is an environment variable that specifies the port for an HTTP service
	EnvHttpPort EnvVariable = "HTTP_PORT"

	// EnvProject is an environment variable to assign a GCP project to assign the service
	EnvProject EnvVariable = "PROJECT"

	// EnvTcpPort is an environment variable that specifies the port for a TCP service
	EnvTcpPort EnvVariable = "TCP_PORT"
)
