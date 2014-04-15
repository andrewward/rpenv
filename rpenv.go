package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var sep = flag.Bool("s", false, "separate arguments with spaces")

func main() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-s ") {
		args := []string{}
		for _, arg := range os.Args {
			args = append(args, strings.Split(arg, " ")...)
		}
		os.Args = args
	}
	flag.Parse()

	envs := os.Environ()

	file, err := os.Open("/etc/rentpath/environment.cfg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "rpenv: %s\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envs = append(envs, scanner.Text())
	}
	file.Close()

	if flag.NArg() == 0 {
		for _, env := range envs {
			line := fmt.Sprintf("%q", env)
			fmt.Println(line[1 : len(line)-1])
		}
		os.Exit(0)
	}

	cmd := exec.Command(flag.Args()[0], flag.Args()[1:]...)
	cmd.Env = envs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if cmd.Process == nil {
		fmt.Fprintf(os.Stderr, "rpenv: %s\n", err)
		os.Exit(1)
	}
	os.Exit(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
}
