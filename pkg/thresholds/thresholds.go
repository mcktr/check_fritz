package thresholds

// GetThresholdsStatus returns true if thresholds are set otherwise false
func GetThresholdsStatus(threshold float64) bool {
	if threshold == -1.0 {
		return false
	}

	return true
}

// CheckLower checks if val is lower than threshold
func CheckLower(threshold float64, val float64) bool {
	if threshold == -1.0 {
		return false
	}

	return threshold > val
}

// CheckUpper checks if val is higher than threshold
func CheckUpper(threshold float64, val float64) bool {
	if threshold == -1.0 {
		return false
	}

	return threshold < val
}
