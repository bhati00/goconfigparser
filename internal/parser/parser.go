package parser

import (
	"fmt"
	"goconfigParser/internal/models"
	"strings"
)

func ParseConfig(fileContent string) (models.Config, error) {
	// intialize the config struct
	config := models.Config{
		Sections: make(map[string]map[string]string),
	}
	var currentSection string
	lines := strings.Split(fileContent, "\n")
	for lineNum, rawLine := range lines {

		if rawLine == "" { // ignore the empty lines
			continue
		}

		if strings.Contains(rawLine, "#") || strings.Contains(rawLine, ";") { // ignore comments
			continue
		}

		// if it's a section header
		if strings.HasPrefix(rawLine, "[") && strings.HasSuffix(rawLine, "]") {
			currentSection = strings.TrimSpace(rawLine[1 : len(rawLine)-1])
			if currentSection == "" {
				return config, fmt.Errorf("section name is empty at line %d", lineNum+1)
			}
			config.Sections[currentSection] = make(map[string]string)
			continue
		}

		// if it's a key-value pair
		if strings.Contains(rawLine, "=") {
			if currentSection == "" {
				return config, fmt.Errorf("key-value pair found before section at line %d", lineNum+1)
			}
			parts := strings.SplitN(rawLine, "=", 2)
			if len(parts) != 2 {
				return config, fmt.Errorf("invalid key-value pair at line %d", lineNum+1)
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "" {
				return config, fmt.Errorf("key is empty at line %d", lineNum+1)
			}
			if value == "" {
				return config, fmt.Errorf("value is empty at line %d", lineNum+1)
			}
			config.Sections[currentSection][key] = value
		} else {
			return config, fmt.Errorf("invalid line format at line %d", lineNum+1)
		}
	}
	return config, nil
}
