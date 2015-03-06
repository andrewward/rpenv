package main

import (
	"os"
	"strings"
)

func rpEnvVars(envUri string) map[string]string {

	rawVars := strings.Split(httpRequestBodyAsString(envUri), "\n")
	rpVars := make(map[string]string)

	for _, kvPair := range rawVars {
		if !strings.HasPrefix(kvPair, "#") && kvPair != "" {
			kvArray := strings.Split(strings.Replace(kvPair, "\"", "", -1), "=")
			rpVars[kvArray[0]] = kvArray[1]
		}
	}
	return rpVars
}

func displayRPEnvVars(env string) {
	printSlice(formattedAsEnvVars(mapAsSortedSlice(updatedEnvs(env))))
}

func updatedEnvs(env string) map[string]string {
	return updateMap(envsMap(), rpEnvVars(envUri(env)))
}

func envUri(env string) string {
	var subdomain string
	switch env {
	case "ci":
		subdomain = "ag-web-01.ci.atl"
	case "qa":
		subdomain = "ag-web-01.qa.atl"
	case "prod":
		subdomain = "ag-web-01.atl"
	case "production":
		subdomain = "ag-web-01.atl"
	default:
		println("Provided environment must be one of 'ci', 'qa', 'prod', or 'production'.")
		os.Exit(1)
	}
	return "http://" + subdomain + ".primedia.com/ops/env/environment.cfg"
}
