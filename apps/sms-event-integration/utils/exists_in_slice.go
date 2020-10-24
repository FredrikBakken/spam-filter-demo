package utils

// ExistsInSlice checks if a given value exists in the given slice
func ExistsInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
