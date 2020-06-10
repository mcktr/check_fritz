package fritzutil

// Contains searches for string in a string slice and returns true if the string is found.
func Contains(s []string, v string) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}
