package leetcode

import "fmt"

const (
	N = 4
	W = 5
)

var (
	wt     = []int{1, 2, 3, 4}
	val    = []int{2, 4, 4, 5}
	maxVal int
)

func main() {
	visit := make(map[int]struct{})
	put(W, 0, visit)
	fmt.Println(maxVal)
}

func put(remainWeight, value int, visit map[int]struct{}) {
	for i := 1; i <= N; i++ {
		if _, ok := visit[i]; ok {
			continue
		}
		if wt[i-1] > remainWeight {
			continue
		}
		visit[i] = struct{}{}
		value += val[i-1]
		put(remainWeight-wt[i-1], value, visit)
		value -= val[i-1]
		delete(visit, i)
	}
	res := 0
	for i := range visit {
		res += val[i-1]
	}
	if res > maxVal {
		maxVal = res
	}
}
