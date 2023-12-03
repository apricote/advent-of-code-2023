package util

func Unique[T comparable](slice []T) []T {
	unique := []T{}
	seen := map[T]bool{}

	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			unique = append(unique, item)
			seen[item] = true
		}
	}

	return unique
}
