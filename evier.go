package main

import (
	"evier/config"
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

	if *verbose {
		fmt.Println("Using command line flags:")
		fmt.Println("  config:", *configPath)
		fmt.Println("  verbose", *verbose)
	}

	if *configPath == "" {
		fmt.Println("Must specify config file")
		os.Exit(1)
	}

	parsedConfig, err := config.ParseFile(*configPath)
	check(err)

	fmt.Println(parsedConfig.Rsync.User)
}
