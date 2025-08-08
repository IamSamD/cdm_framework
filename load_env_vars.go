package cdm_framework

import (
	"fmt"
	"os"
)

func LoadEnvVars(envvars []string) (map[string]string, error) {
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
