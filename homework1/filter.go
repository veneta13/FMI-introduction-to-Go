package main

func Filter(p func(int) bool) func(...int) []int {
	return func(args ...int) []int {
		var answer []int
		for _, item := range args {
			if p(item) {
				answer = append(answer, item)
			}
		}
		return answer
	}
}
