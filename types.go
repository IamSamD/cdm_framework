package cdm_framework

// Config is the config map used to access all config values within a checks logic code.
type Config map[string]string

// Runfunc is the function type for the callback function accepted by RunCheck
type RunFunc func(Config) error
