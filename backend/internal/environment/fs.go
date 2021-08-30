package environment

import (
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func (env *Environment) ApplyFirestoreOpts(opts *[]option.ClientOption) error {
	if !env.IsLocal() {
		return nil
	}

	if !env.FirestoreEmulatorHost.Exists() {
		return nil
	}

	conn, err := grpc.Dial(env.FirestoreEmulatorHost.String(), grpc.WithInsecure())
	if err != nil {
		return err
	}

	*opts = append(*opts, option.WithGRPCConn(conn))
	return nil
}
