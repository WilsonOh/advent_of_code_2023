package aoc

import "strconv"

func Map[T any, U any](l []T, f func(t T) U) []U {
	ret := make([]U, len(l))
	for i, a := range l {
		ret[i] = f(a)
	}
	return ret
}

func Filter[T any](l []T, f func(t T) bool) []T {
	ret := make([]T, 0)
	for _, a := range l {
		if f(a) {
			ret = append(ret, a)
		}
	}
	return ret
}

func Reduce[T any, U any](l []T, reducerFunc func(curr T, acc U) U, init U) U {
	ret := init
	for _, a := range l {
		ret = reducerFunc(a, ret)

	}
	return ret
}

func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func MapStrToInt(l []string) []int {
	return Map(l, StrToInt)
}
