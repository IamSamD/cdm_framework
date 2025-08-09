package cdm_framework

func RunCheck(config Config, runFunc RunFunc) error {
	if err := runFunc(config); err != nil {
		return err
	}

	return nil
}
