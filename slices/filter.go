package slices

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Filter[T constraints.Ordered](items []T, callback func(val T) bool) []T {
	result := make([]T, 0, len(items))

	for _, item := range items {
		if callback(item) {
			result = append(result, item)
		}
	}

	log(fmt.Sprintf("%s the result is %v", pkgName, result))
	return result
}
