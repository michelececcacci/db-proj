package util

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

func Min[T constraints.Ordered](x, y T) T {
	if x <= y {
		return x
	}
	return y
}
