package sort

import (
	"GoAlgo/myUtils"
	"fmt"
)

func boubleSort(comparefunc func(a interface{}, b interface{}) bool, vals []interface{}) []interface{} {
	move := false
	for i := len(vals) - 1; i > 0; i-- {
		move = false
		for j := 0; j < i; j++ {
			if comparefunc(vals[j], vals[j+1]) {
				vals[j], vals[j+1] = vals[j+1], vals[j]
				move = true
			}
		}
		if !move {
			break
		}
	}
	return vals
}

func Run() {
	numbers := myUtils.ToInterfaceArray(1, 7, 2, 4, 9, 0, 4, 5)

	fmt.Printf("before sort: %v\n", numbers)
	out := boubleSort(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	}, numbers)
	fmt.Printf("after sort: %v\n", out)
}
