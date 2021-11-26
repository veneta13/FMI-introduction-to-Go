package main

func Reducer(initial int, f func(int, int) int) func(...int) int {
	count := initial
	return func(args ...int) int {
		for _, item := range args {
			count = f(count, item)
		}
		return count
	}
}
