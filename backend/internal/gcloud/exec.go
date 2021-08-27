package gcloud

import (
	"bytes"
	"os/exec"
)

var (
	Config = &GCloudConfig{}
)

const (
	name = "gcloud"
)

type GCloudConfig struct {
}

func which() (string, error) {
	return exec.LookPath(name)
}

func cmd(args ...string) (*exec.Cmd, error) {
	path, err := which()
	if err != nil {
		return nil, err
	}
	return exec.Command(path, args...), nil
}

func (config *GCloudConfig) Project() (string, error) {
	c, err := cmd("config", "get-value", "project")
	if err != nil {
		return "", err
	}
	b, err := c.Output()
	if err != nil {
		return "", err
	}
	b = bytes.TrimSpace(b)
	return string(b), nil
}
