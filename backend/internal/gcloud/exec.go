package gcloud

import (
	"bytes"
	"os/exec"
)

const (
	name = "gcloud"
)

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

func output(args ...string) (string, error) {
	c, err := cmd(args...)
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
