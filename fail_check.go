package cdm_framework

import "os"

// FailCheck triggers the failure of a check
// It simple uses os.Exit(1) to exit the check with an exit code of 1
// It is included in the cdm_framework in order to ensure that all checks fail with the same behaviour while also providing a more high level experience for check development
func FailCheck() {
	os.Exit(1)
}
