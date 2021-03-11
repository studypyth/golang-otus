package main

import (
	"errors"
	"os"
	"os/exec"
)

func RunCmd(commands []string, env Environment) (returnCode int) {
	code := 0
	for k, v := range env {
		if !v.NeedRemove && v.Value != "" {
			os.Setenv(k, v.Value)
		} else {
			os.Unsetenv(k)
		}
	}
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			code = exitError.ExitCode()
		}
	}
	return code
}
