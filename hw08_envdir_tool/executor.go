package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
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
		if exitError, ok := err.(*exec.ExitError); ok {
			code = exitError.ExitCode()
		}
	}
	return code
}
