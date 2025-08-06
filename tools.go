package main

import (
	"context"
	"github.com/moby/moby/api/types/container"
	Docker "github.com/moby/moby/client"
	"io"
)

// the available tools

type Tool interface {
	Run() string
	Call() string
	String() string
}

type ContainerCommandRun struct {
	Command      string
	DockerClient Docker.Client
}

func (u *ContainerCommandRun) Run() string {
	var output string
	ctx := context.Background()
	resp, err := u.DockerClient.ContainerCreate(ctx, &container.Config{
		Image: "python-dev",
		Cmd:   []string{"bash -c", u.Command},
	}, nil, nil, nil, "")
	if err != nil {
		output = "error " + err.Error()
	}

	options := container.LogsOptions{ShowStdout: true}
	out, err := u.DockerClient.ContainerLogs(ctx, resp.ID, options)
	if err != nil {
		output = "error " + err.Error()
	}
	data, err := io.ReadAll(out)
	output = string(data)

	return output
}

func (u *ContainerCommandRun) Call() string {
	return u.Run()
}

func (u *ContainerCommandRun) String() string {

	return ""
}
