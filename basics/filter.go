package main

import "fmt"

func filter(pred func(int) bool, values []int) []int {
	var answer []int
	for _, v := range values {
		if pred(v) {
			answer = append(answer, v)
		}
	}

	return answer
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(filter(isOdd, values)) // [1 3 5 7]
}
