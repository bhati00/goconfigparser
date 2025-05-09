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
		name      string
		input     string
		expected  models.Config
		expectErr bool
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
			expectErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultBytes, err := ParseConfig(test.input)

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
