package generics

func identity[T any](in T) T {
	return in
}
