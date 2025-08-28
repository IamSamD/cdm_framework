package cdm_framework

import "fmt"

func CheckFailedReport(message string) {
	fmt.Print("\n\n====  CHECK FAILED ====\n\n")
	fmt.Print(message + "\n\n")
}
