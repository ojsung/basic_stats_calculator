package slice_utils

func Contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func Map[T any, U any](input[]T, mapFunction func(T) U) (output []U) {
	output = make([]U, len(input))
	for index, value := range input {
		output[index] = mapFunction(value)
	}
	return output
}