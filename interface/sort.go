package main

import (
	"fmt"
	"sort"
)

type Sort struct {
	m []int
}

func (s Sort) Len() int {
	return len(s.m)
}

// 表示使用什么标准进行排序
func (s Sort) Less(i int, j int) bool {
	return s.m[i] > s.m[j]
}

func (s Sort) Swap(i, j int) {
	s.m[i], s.m[j] = s.m[j], s.m[i]
}
func main() {

	// 都是升序
	sint := []int{3,52,532,51,2}
	sort.Ints(sint)
	sfloat := []float64{3,52.4,532,52.1,2}
	sort.Float64s(sfloat)

	strings := []string{"dqfq", "klej", "dwqh", "aaa", "中国", "bbb"}
	sort.Strings(strings)
	s := []int{3,521,532,52,2, 52}
	sort.Ints(s)

	// 二分查找
	index := sort.Search(len(s), func(i int) bool {
	return s[i] == 53
	})
	fmt.Println(s, index)

	s1 := Sort{m: []int{3,52,532,51,2}}
	sort.Sort(s1)
	fmt.Println(s1)
}