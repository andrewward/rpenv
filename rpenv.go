package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()

	cmdStatus := 0

	if flag.NArg() == 0 {
		println("must provide an environment, e.g. 'ci', 'qa', or 'prod'...")
	} else if flag.NArg() == 1 {
		displayRPEnvVars(flag.Args()[0])
	} else {
		envVars := formattedAsEnvVars(mapAsSortedSlice(updatedEnvs(flag.Args()[0])))
		cmdStatus = executeCommand(flag.Args(), envVars)
	}

	os.Exit(cmdStatus)
}
