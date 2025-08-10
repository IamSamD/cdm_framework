package cdm_framework

// RunCheck is a higher order function used to centralise the behaviour of running a check
// It accepts the config object and the Check callback function which holds the check logic
// It will return any errors generated in the check logic
func RunCheck(config Config, runFunc RunFunc) error {
	if err := runFunc(config); err != nil {
		return err
	}

	return nil
}
