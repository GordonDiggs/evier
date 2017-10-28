package main

import (
	"evier/config"
	"evier/runner"
	"evier/version"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	configPath := flag.String("config", "", "Config path")
	verbose := flag.Bool("verbose", false, "Verbose mode")
	versionPtr := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *versionPtr {
		fmt.Printf("evier v%s\n", version.Version)
		os.Exit(0)
	}

	if *configPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *verbose {
		fmt.Println("Using command line flags:")
		fmt.Println("  config:", *configPath)
		fmt.Println("  verbose", *verbose)
	}

	parsedConfig, err := config.ParseFile(*configPath)
	check(err)

	err = runner.Run(parsedConfig)
	check(err)
}
