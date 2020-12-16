package main

func sum(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}
	return slice[0] + sum(slice[1:])
}

func count(slice []int) int {
	var i = 1
	if len(slice) == 1 {
		return 1
	}
	return i + count(slice[1:])
}

func max(slice []int) int {
	var maxim int
	if len(slice) == 2 {
		if slice[0] > slice[1] {
			return slice[0]
		} else {
			return slice[1]
		}
	}
	maxim = max(slice[1:])
	if slice[0] > maxim {
		return slice[0]
	} else {
		return maxim
	}
}
