package main

import (
	"flag"
	"fmt"
	"goconfigparser/internal/parser" // Assuming the parser package exists in your project
	"os"
)

func main() {
	configPath := flag.String("input", "", "Path to the configuration file")
	outPutPath := flag.String("output", "", "Path to the output file")
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

	outPutDir := "output"
	if *outPutPath != "" {
		err = os.MkdirAll(outPutDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}
	fullPath := fmt.Sprintf("%s/%s", outPutDir, *outPutPath)
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Printf("File %s already exists. Overwrite? (y/n): ", fullPath)
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("Operation cancelled by user.")
			os.Exit(0)
		}
	}

	err = os.WriteFile(fullPath, []byte(config), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("File %s created successfully\n", fullPath)
	}
	fmt.Printf("Parsed config %s\n", config)
}
