package main

import (
	"context"
	"fmt"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/testdata"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"strconv"
)

var (
	rootCmd = &cobra.Command{
		Use:   "client",
		Short: "Call gRPC counter service",
	}

	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List counts",
		PreRunE: preRun,
		RunE:    list,
	}

	updateCmd = &cobra.Command{
		Use:     "update [COUNT]",
		Short:   "Update user count with COUNT (COUNT >= 0; default: 0)",
		PreRunE: preRun,
		RunE:    update,
		Args: cobra.MaximumNArgs(1),
	}

	streamCmd = &cobra.Command{
		Use:     "stream",
		Short:   "Stream counts",
		PreRunE: preRun,
		RunE:    stream,
	}

	endpoint      string
	limit int64
	opts          []grpc.DialOption
	counterClient counter.CounterServiceClient
)

func init() {
	env()
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", fmt.Sprintf("localhost:%s", environment.DefaultEnvironment.PortTcp.String()), "Server endpoint")
	rootCmd.AddCommand(listCmd, updateCmd, streamCmd)
	listCmd.Flags().Int64VarP(&limit, "limit", "n", -1, "Limit results (<0 is no limit)")
}

func env() {
	if err := testdata.LoadDotEnv(); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func list(c *cobra.Command, _ []string) error {
	ctx := c.Context()
	resp, err := counterClient.ListCounts(ctx, &counter.ListCountsRequest{
		Limit: limit,
	})
	if err != nil {
		return err
	}
	for _, item := range resp.Items {
		fmt.Println(item.String())
	}
	return nil
}

func update(c *cobra.Command, args []string) error {
	ctx := c.Context()
	var count int64 = 0
	if len(args) > 0 {
		v, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}
		count = v
	}

	resp, err := counterClient.UpdateCount(ctx, &counter.UpdateCountRequest{
		Count: count,
	})

	if err != nil {
		return err
	}

	fmt.Printf("updated with %v\n", resp.Count)

	return nil
}

func stream(c *cobra.Command, _ []string) error {
	ctx := c.Context()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	client, err := counterClient.StreamCounts(ctx, &counter.StreamCountsRequest{})
	if err != nil {
		return err
	}
	errChan := make(chan error)
	go func(ctx context.Context) {
		for {
			count, err := client.Recv()
			if err != nil {
				errChan <- err
			}
			if count == nil {
				continue
			}
			fmt.Println(count.String())
		}
	}(ctx)
	for {
		select {
		case err := <- errChan:
			return err
		case <- ctx.Done():
			return nil
		}
	}
}
