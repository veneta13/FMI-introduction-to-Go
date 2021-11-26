package main

func MapReducer(initial int, mapper func(int) int, reducer func(int, int) int) func(...int) int {
	count := initial
	return func(args ...int) int {
		var mapArray []int
		for _, item := range args {
			mapArray = append(mapArray, mapper(item))
		}

		for _, item := range mapArray {
			count = reducer(count, item)
		}
		return count
	}
}
