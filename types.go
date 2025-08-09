package cdm_framework

type Config map[string]string

type RunFunc func(Config) error
