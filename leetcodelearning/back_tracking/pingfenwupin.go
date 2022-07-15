package back_tracking

import (
	"fmt"
	"math"
)

func main() {
	var exampleCount int
	fmt.Scan(&exampleCount)
	for exampleCount > 0 {
		var n int
		fmt.Scan(&n)
		things := make([]int, 0)
		for n > 0 {
			var value int
			fmt.Scan(&value)
			things = append(things, value)
			n--
		}
		fmt.Println(f(things))
		exampleCount--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(things []int) int {
	res := math.MaxInt
	var dfs func(layer, value1, value2, weight int)
	dfs = func(layer, value1, value2, weight int) {
		if layer == len(things) {
			if value1 == value2 {
				res = min(weight, res)
			}
			return
		}
		// 分配给第一个人
		dfs(layer+1, value1+things[layer], value2, weight)
		// 分配给第二个人
		dfs(layer+1, value1, value2+things[layer], weight)
		//	丢掉
		dfs(layer+1, value1, value2, weight+things[layer])
	}
	dfs(0, 0, 0, 0)
	return res
}
