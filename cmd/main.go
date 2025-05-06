package main

import (
	"flag"
	"fmt"
	"goconfigparser/internal/parser" // Assuming the parser package exists in your project
	"os"
)

func main() {
	configPath := flag.String("file", "", "Path to the configuration file")
	flag.Parse()
	if *configPath == "" {
		fmt.Println("Please provide a configuration file path using -file flag")
		os.Exit(1)
	}
	content, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	config, err := parser.ParseConfig(string(content))
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parsed config %s\n", config)
}
