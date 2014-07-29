package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func executeCommand(c []string, envs []string) int {
	exitStatus := 0
	cmd := exec.Command(c[1], c[2:]...)
	cmd.Env = envs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if cmd.Process == nil {
		fmt.Fprintf(os.Stderr, "rpenv: %s\n", err)
		exitStatus = 1
	}

	exitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()

	return exitStatus
}
