package helper

import "os"

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, its value is returned.
// Otherwise, the provided default value is returned.
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
