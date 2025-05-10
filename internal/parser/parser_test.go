// testing the  parser functions for cases, valid configuration, empty configuration, empty section, invalid line
package parser

import (
	"encoding/json"
	"goconfigparser/internal/models"
	"reflect"
	"testing"
)

func TestParsingConfig(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    models.Config
		expectError bool
	}{
		{
			name: "valid config",
			input: `
			[server]
			host = localhost
			port = 8080

			[database]
			user = root
			password = secret
			`,
			expected: models.Config{
				Sections: map[string]map[string]string{
					"server": {
						"host": "localhost",
						"port": "8080",
					},
					"database": {
						"user":     "root",
						"password": "secret",
					},
				},
			},
			expectError: false,
		},
		{
			name: "Empty config",
			input: `
			
			`,
			expected:    models.Config{Sections: map[string]map[string]string{}},
			expectError: false,
		},
		{
			name: "Comment and whitespace only",
			input: `
			; this is a comment
			# another comment

			`,
			expected:    models.Config{Sections: map[string]map[string]string{}},
			expectError: false,
		},
		{
			name: "Key-value pair before section",
			input: `
			key = value
			`,
			expectError: true,
		},
		{
			name: "Section with no key-values",
			input: `
			[section]
			`,
			expected: models.Config{
				Sections: map[string]map[string]string{
					"section": {},
				},
			},
			expectError: false,
		},
		{
			name: "Empty section name",
			input: `
			[]
			key=value
			`,
			expectError: true,
		},
		{
			name: "Empty key",
			input: `
			[section]
			= value
			`,
			expectError: true,
		},
		{
			name: "Empty value",
			input: `
			[section]
			key =
			`,
			expectError: true,
		},
		{
			name: "Invalid line format",
			input: `
			[section]
			justsomegarbage
			`,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultBytes, err := ParseConfig(test.input)
			if test.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
				return // no need to continue
			}

			if err != nil {
				t.Fatalf("Did not expect error but got: %v", err)
			}
			var actual models.Config
			if err := json.Unmarshal(resultBytes, &actual); err != nil {
				t.Fatalf("Failed to unmarshal result: %v", err)
			}

			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf("Expected %+v, got %+v", test.expected, actual)
			}
		})
	}
}
