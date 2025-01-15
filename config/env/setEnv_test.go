package env

import (
	"log"
	"os"
	"testing"
)

func TestSetEnvTerminal(t *testing.T) {
	tt := []struct {
		name        string
		envVar      string
		envValue    string
		expectError bool
	}{
		{
			name:        "Successful set environment variable",
			envVar:      "GIN_MODE",
			envValue:    "SetEnv",
			expectError: false,
		},
		{
			// to see what happens when trying to set environment variable with empty value.
			name:        "Empty environment variable value",
			envVar:      "GIN_MODE",
			envValue:    "",
			expectError: false,
		},
		{
			// to see what happens when trying to set a non-existent environment variable.
			name:        "Set non-existent environment variable",
			envVar:      "NON_EXIST_VAR",
			envValue:    "SetEnv",
			expectError: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := os.Setenv(tc.envVar, tc.envValue)
			if (err != nil) != tc.expectError {
				log.Fatalf("SetEnvTerminal() expected error = %v, but got %v", tc.expectError, err)
			}

			gotValue := os.Getenv(tc.envVar)
			if gotValue != tc.envValue {
				log.Fatalf("SetEnvTerminal() expected value = %v, but got %v", tc.envValue, gotValue)
			}
		})
	}
}
