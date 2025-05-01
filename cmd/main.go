package main

import (
	"fmt"
	"goconfigparser/internal/parser" // Assuming the parser package exists in your project
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: configparser <config_file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	config, err := parser.ParseConfig(string(content))
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parsed config: %+v\n", config)
}
