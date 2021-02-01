package utils

import "strings"

//StringTreatment string treatment
func StringTreatment(from, to string) (string, string) {
	return strings.ToUpper(strings.TrimSpace(from)), strings.ToUpper(strings.TrimSpace(to))
}
