package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "undefined"

func printVersion() {
	fmt.Println("{{ cookiecutter.repo_name }} version:", version)
	os.Exit(0)
}

// Documentation for main function
func main() {
	showVersion := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *showVersion {
		printVersion()
	}
}
