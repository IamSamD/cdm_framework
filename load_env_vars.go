package cdm_framework

import (
	"fmt"
	"os"
)

// used as part of InitEnvironment to load env vars and parse the config map
func loadEnvVars(envvars []string) (map[string]string, error) {
	envVarMap := make(map[string]string)

	for _, envVar := range envvars {
		value, exists := os.LookupEnv(envVar)
		if !exists {
			return nil, fmt.Errorf("environment variable %s not set", envVar)
		}
		envVarMap[envVar] = value
	}
	return envVarMap, nil
}
