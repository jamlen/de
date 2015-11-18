package main

import (
    "bytes"
    "errors"
	"os/exec"
    "strings"
)

type ShellRunner struct {
}

func (r *ShellRunner) Run(command string, args []string) error {
	var stdout bytes.Buffer
	cmd := exec.Command(command, strings.Join(args, " "))
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		// If "getent" is missing, ignore it
		if err != exec.ErrNotFound {
			return err
		}
	}
	result := strings.TrimSpace(stdout.String())
    if result != "" {
        return errors.New("Booom")
    }
    return nil
}
