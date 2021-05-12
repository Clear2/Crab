package main

import "math"

type Employee struct {
	Happy int  // 这名员工可以带来的快乐值
	List []Employee	// 这名员工有哪些直接下级
}

type NeedInfo struct {
	 LaiMaxHappy int
	 buMaxHappy int
}

func processE(x Employee) *NeedInfo{
	if x.List == nil {
		return &NeedInfo{LaiMaxHappy: x.Happy, buMaxHappy: 0}
	}
	lai := x.Happy  // x 来的情况下
	var bu float64
	for _, e := range x.List {
		nextInfo := processE(e)
		lai += nextInfo.buMaxHappy
		bu += math.Max(float64(nextInfo.LaiMaxHappy), float64(nextInfo.buMaxHappy))
	}
	return &NeedInfo{LaiMaxHappy: lai, buMaxHappy: int(bu)}
}
func main() {
	//n2 := &Employee{Happy: 20}
	//n1 := &Employee{Happy: 20}
	//root := &Employee{Happy: 10, List: }

}