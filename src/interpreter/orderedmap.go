package interpreter

// a key can only be a string or an integer
type Keys interface {
	int | string
}

type OrderedMap[K Keys, V any] struct {
	Items map[K]*Element[K, V]
	// LinkedList List
}

type Element[K comparable, V any] struct {
	Key   K
	Value V
}
