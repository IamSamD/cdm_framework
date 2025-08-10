package cdm_framework

import (
	"github.com/joho/godotenv"
)

// InitEnvironment reads and parses environment variables which act as confguratio values for the check
// It ustilises godotenv so that env vars can easily be fed in for local testing if desired
// It accepts a slice of strings which are env var names defined in the check code by the developer of the check
// It returns a config map with values mapped to names for easy access in check logic
// The function is utilised in the initialisation of a check and check developers should not need to use it
func InitEnvironment(configValues []string) (Config, error) {
	var config Config
	// Use godotenv for testing locally
	goDotEnvConfig, err := godotenv.Read()
	if err != nil {
		configValuesMap, err := loadEnvVars(configValues)
		if err != nil {
			return nil, err
		}

		config = configValuesMap
	}
	config = goDotEnvConfig

	return config, nil
}
