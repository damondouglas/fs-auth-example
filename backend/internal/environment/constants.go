package environment

import (
	"flag"
	"fmt"
	"os"
)

var (
	EnvHttpPort EnvVariable = "HTTP_PORT"
	EnvTcpPort  EnvVariable = "TCP_PORT"
	FlagLocal               = &FlagVariable{
		Flag: &flag.Flag{
			Name:  "local",
			Usage: "Flag as local environment",
		},
	}

	FlagVerbose = &FlagVariable{
		Flag: &flag.Flag{
			Name:  "verbose",
			Usage: "Turn on verbose logging",
		},
	}

	Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println()
		fmt.Println("A gRPC service that manages user counts.")
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Environment variables:")
		for _, v := range []EnvVariable{
			EnvHttpPort,
			EnvTcpPort,
		} {
			fmt.Println(v.Key())
		}
	}
)

func init() {
	FlagLocal.value = flag.Bool(FlagLocal.Name, false, FlagLocal.Usage)
	FlagVerbose.value = flag.Bool(FlagVerbose.Name, false, FlagVerbose.Usage)
}
