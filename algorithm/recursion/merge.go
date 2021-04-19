package main

func main() {
	arr := []int{10, 3, 9, 18, 5, 1}
	process1(arr, 0, len(arr)-1)
}
func process1(arr []int, L, R int) {

}

func merge(arr []int, L, M, R int) {
	help := make([]int, L+R-1)
	i, p1, p2 := 0, L, M+1
	for p1 <= M && p2 <= R {
		i++
		if arr[p1] <= arr[p2] {
			p1++
			help[i] = arr[p1]
		} else {
			help[i] = arr[p2]
		}
	}
	for p1 <= M {
		i++
		p1++
		help[i] = arr[p1]
	}
	for p2 <= R {
		i++
		p2++
		help[i] = arr[p2]
	}
	for i = 0; i< len(help); i++ {
		arr[L + i ] = help[i]
	}

}
