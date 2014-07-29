package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "rpenv: %s\n", err)
		os.Exit(1)
	}
}

func httpRequestBodyAsString(uri string) string {
	resp, err := http.Get(uri)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return string(body)
}

func printSlice(s []string) {
	println(strings.Join(s, "\n"))
}
