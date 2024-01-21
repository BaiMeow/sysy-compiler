package utils

func Map[T any, V any](s []T, f func(T) V) []V {
	if s == nil {
		return nil
	}
	vs := make([]V, len(s))
	for i, v := range s {
		vs[i] = f(v)
	}
	return vs
}
