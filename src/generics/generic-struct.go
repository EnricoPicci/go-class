package generics

type node[T any] struct {
	data        T
	letf, right *node[T]
}
