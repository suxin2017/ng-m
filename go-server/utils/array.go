package utils

func Contains[T interface{}](elems []T, v func(element T) bool) bool {
	for _, s := range elems {
		if v(s) {
			return true
		}
	}
	return false
}
