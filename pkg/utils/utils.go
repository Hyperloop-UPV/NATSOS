package utils

// Contains checks if a slice contains a specific value. It is a generic function that works with any comparable type.
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
