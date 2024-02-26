package slices

import "fmt"

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			log(fmt.Sprintf("%s found %v into the list", pkgName, value))
			return true
		}
	}
	log(fmt.Sprintf("%s not found %v into the list", pkgName, value))
	return false
}
