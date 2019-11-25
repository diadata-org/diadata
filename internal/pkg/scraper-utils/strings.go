package utils

// Contains takes a slice of strings and a string and checks if it is contained in the slice.
func Contains(s *[]string, str string) bool {
	for _, a := range *s {
		if a == str {
			return true
		}
	}
	return false
}
