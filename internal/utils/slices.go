package utils

func ForEach[T any](items []T, action func(item T)) {
	for _, item := range items {
		action(item)
	}
}
