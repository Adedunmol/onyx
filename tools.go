package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/moby/moby/api/types/container"
	Docker "github.com/moby/moby/client"
	"io"
	"os"
	"strings"
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

func (c *ContainerCommandRun) Run() string {
	var output string
	ctx := context.Background()
	resp, err := c.DockerClient.ContainerCreate(ctx, &container.Config{
		Image: "python-dev",
		Cmd:   []string{"bash -c", c.Command},
	}, nil, nil, nil, "")
	if err != nil {
		output = "error " + err.Error()
	}

	options := container.LogsOptions{ShowStdout: true}
	out, err := c.DockerClient.ContainerLogs(ctx, resp.ID, options)
	if err != nil {
		output = "error " + err.Error()
	}
	data, err := io.ReadAll(out)
	output = string(data)

	return output
}

func (c *ContainerCommandRun) Call() string {
	return c.Run()
}

func (c *ContainerCommandRun) String() string {

	return ""
}

type UpsertFile struct {
	FilePath, content string
	DockerClient      Docker.Client
}

func (u *UpsertFile) Run() string {
	ctx := context.Background()

	cmd := fmt.Sprintf("sh -c \"cat > %s\"", u.FilePath)

	resp, err := u.DockerClient.ContainerCreate(ctx, &container.Config{
		Image: "python-dev",
		Cmd:   []string{cmd},
	}, nil, nil, nil, "")
	if err != nil {
		return "error " + err.Error()
	}

	execResp, err := u.DockerClient.ContainerExecCreate(ctx, resp.ID, container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	})
	if err != nil {
		return "error " + err.Error()
	}

	attachResp, err := u.DockerClient.ContainerExecAttach(ctx, execResp.ID, container.ExecAttachOptions{
		Tty: true,
	})
	if err != nil {
		return "error " + err.Error()
	}
	defer attachResp.Close()

	scanner := bufio.NewScanner(strings.NewReader(u.content))
	for scanner.Scan() {
		_, err := fmt.Fprintln(attachResp.Conn, scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "write error:", err)
			break
		}
	}

	return "file written successfully"
}

func (u *UpsertFile) Call() string {
	return u.Run()
}

func (u *UpsertFile) String() string {

	return ""
}
