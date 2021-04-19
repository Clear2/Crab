package main

import "fmt"

func main() {
	arr := []int{100, 97, 32, 165, 40}
	fmt.Println(getMax(arr))
}

func getMax(arr[] int) int {
	return process(arr, 0, len(arr) - 1)
}
func process(arr []int, L int, R int)  int {
	if L == R {
		return arr[L]
	}
	mid := L + ((R - L ) >>  1)

	leftMax := process(arr, L, mid)
	rightMax := process(arr, mid + 1, R)
	if leftMax < rightMax {
		return rightMax
	}
	return leftMax
}