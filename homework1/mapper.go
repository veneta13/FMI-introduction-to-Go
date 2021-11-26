package main

func Mapper(f func(int) int) func(...int) []int {
	return func(args ...int) []int {
		var answer []int
		for _, item := range args {
			answer = append(answer, f(item))
		}
		return answer
	}
}
