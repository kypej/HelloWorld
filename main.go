package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

//за наименьшее число шагов найти максимальное значение
var b = []int{88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87}

func main() {
	http.HandleFunc("/first", first)
	http.HandleFunc("/second", second)
	fmt.Println(sum(a))
	fmt.Println(count(a))
	fmt.Println(max(a))
	fmt.Println(maxShift(b))
	//http.ListenAndServe(":7777", nil)
}

func maxShift(slice []int) int {
	i := len(slice) - slice[len(slice)-1]
	return slice[i-1]
}

func binarSearch(number int) int {
	high := len(a) - 1
	low := 0
	var mid int
	for high >= low {
		mid = (high + low) / 2
		n := a[mid]
		if n == number {
			return mid
		}
		if n > number {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search(number int) int {
	for i := range a {
		if a[i] == number {
			return i
		}
	}
	return -1
}

func first(w http.ResponseWriter, req *http.Request) {
	n, _ := strconv.Atoi(req.Header.Get("n"))
	var a int
	for i := 0; i < 1000000000; i++ {
		a = binarSearch(n)
	}
	w.Write([]byte(strconv.Itoa(a)))
}

func second(w http.ResponseWriter, req *http.Request) {
	n, _ := strconv.Atoi(req.Header.Get("n"))
	var a int
	for i := 0; i < 1000000000; i++ {
		a = search(n)
	}
	w.Write([]byte(strconv.Itoa(a)))
}
