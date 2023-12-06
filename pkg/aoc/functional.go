package aoc

import "strconv"

func Map[T any, U any](l []T, f func(t T) U) []U {
	ret := make([]U, len(l))
	for i, a := range l {
		ret[i] = f(a)
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
