package utils

func ForEach[T any](items []T, action func(item T)) {
	for _, item := range items {
		action(item)
	}
}

func TransformSlice[X any, Z any](input []X, transformer func(X) Z) []Z {
	output := make([]Z, len(input))
	for i, x := range input {
		output[i] = transformer(x)
	}
	return output
}
