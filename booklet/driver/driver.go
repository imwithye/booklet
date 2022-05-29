package driver

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Pull(ctxs ...context.Context) error {
	var ctx context.Context
	if len(ctxs) == 0 || ctxs[0] == nil {
		ctx = context.Background()
	} else {
		ctx = ctxs[0]
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	io.Copy(os.Stdout, reader)
	return err
}

func Compile(ctxs ...context.Context) error {
	var ctx context.Context
	if len(ctxs) == 0 || ctxs[0] == nil {
		ctx = context.Background()
	} else {
		ctx = ctxs[0]
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)

	return nil
}
