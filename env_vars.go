package main

import (
	"os"
	"strings"
)

func envsMap() map[string]string {
	envsMap := make(map[string]string)

	for _, kvPair := range os.Environ() {
		kvArray := strings.Split(kvPair, "=")
		envsMap[kvArray[0]] = kvArray[1]
	}

	return envsMap
}

func formattedAsEnvVars(s [][2]string) []string {
	formatted := make([]string, 0)
	for _, kvPair := range s {
		formatted = append(formatted, kvPair[0]+"="+kvPair[1])
	}
	return formatted
}
