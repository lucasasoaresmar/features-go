package helpers

import (
	"os"
)

// EnvOrDefault returns a env value or the thefault option
func EnvOrDefault(envName string, defaultValue string) string {
	env, ok := os.LookupEnv(envName)
	if ok {
		return env
	}
	return defaultValue
}
