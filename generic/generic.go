package genericutils

func ZeroOf[T any]() (_ T) { return }

func Ptr[T any](v T) *T {
	return &v
}

func NilEmpty[T comparable](v T) *T {
	if v == ZeroOf[T]() {
		return nil
	}
	return Ptr(v)
}

func Coalesce[T any](v *T, def T) T {
	if NilEmpty(v) == nil {
		return def
	}
	return *v
}

func CoalesceMapString[V any](m map[string]V, key string, def V) V {
	value, ok := m[key]
	if !ok {
		return def
	}

	if NilEmpty(&value) == nil {
		return def
	}

	return value
}

func EmptyIfNil[T comparable](v *T) T {
	if v == nil {
		return ZeroOf[T]()
	}
	return *v
}
