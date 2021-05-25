package main

import "fmt"

/// 冒泡 插入
// 0 - N-1 比较N次，才确定谁放在0，才搞定了一个树
// 比较 N-2次， 又搞定了一个树
// 时间复杂度 N*LogN
// 浪费比较行为

// 归并排序，没有浪费比较行为
func mergeSort(s *[]int) {
	if len(*s) < 2 {
		return
	}
	merge(s, 0, len(*s) -1)
}

func merge(arr *[]int, L, R int)  {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	merge(arr, L, mid)
	fmt.Println(arr)
	merge(arr, mid+1, R)
	//mergeProcess(arr, L, mid, L)
}

func mergeProcess(arr []int, L, mid, R int)  {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid+1
	for p1 <= mid && p2 <= R{
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
		} else {
			help[i] = arr[p2]
		}
		i+=1
		p1+=1
		p2+=1
	}

	for p1 <= mid {
		help[i] = arr[p1]
		i+=1
		p1+=1
	}
	for p2 <= R {
		help[i] = arr[p2]
		i+=1
		p2+=1
	}
	for j := 0; j < len(help); j++ {
		arr[L + j] = help[j]
	}
}

// T(N) =
func main() {
	//  78, 32, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106
	a := []int{2, 1, 3, 4, 50}
	mergeSort(&a)
}
