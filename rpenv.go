package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var sep = flag.Bool("s", false, "separate arguments with spaces")

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "rpenv: %s\n", err)
		os.Exit(1)
	}
}

func envNotProvidedExit() {
	println("must provide an environment, e.g. 'ci', 'qa', or 'prod'...")
	os.Exit(1)
}

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

	var envFile string

	if len(flag.Args()) == 0 {
		envNotProvidedExit()
	}

	appEnv := flag.Args()[0]

	switch appEnv {
	case "ci":
		envFile = "http://ag-web-01.ci.nor.primedia.com/ops/env/environment.cfg"
	case "qa":
		envFile = "http://ag-web-01.qa.nor.primedia.com/ops/env/environment.cfg"
	case "prod":
		envFile = "http://ag-web-01.atl.primedia.com/ops/env/environment.cfg"
	case "production":
		envFile = "http://ag-web-01.atl.primedia.com/ops/env/environment.cfg"
	default:
		envNotProvidedExit()
	}

	resp, err := http.Get(envFile)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	envArray := strings.Split(string(body), "\n")
	for _, kvPair := range envArray {
		if !strings.HasPrefix(kvPair, "#") && kvPair != "" {
			envs = append(envs, strings.Replace(kvPair, "\"", "", -1))
		}
	}

	if flag.NArg() == 0 {
		for _, env := range envs {
			line := fmt.Sprintf("%q", env)
			fmt.Println(line[1 : len(line)-1])
		}
		os.Exit(0)
	}

	cmd := exec.Command(flag.Args()[1], flag.Args()[2:]...)
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
