package cdm_framework

import (
	"github.com/joho/godotenv"
)

func InitEnvironment(configValues []string) (Config, error) {
	var config Config
	// Use godotenv for testing locally
	goDotEnvConfig, err := godotenv.Read()
	if err != nil {
		configValuesMap, err := LoadEnvVars(configValues)
		if err != nil {
			return nil, err
		}

		config = configValuesMap
	}
	config = goDotEnvConfig

	return config, nil
}
