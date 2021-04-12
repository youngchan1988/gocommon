package sliceutils

import (
	"sort"
)

func Ints(a []int) {
	sort.Sort(IntSlice(a))
}

func RInts(a []int) {
	sort.Sort(sort.Reverse(IntSlice(a)))
	// 等价于
	//sort.Slice(IntSlice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Int64s(a []int64) {
	sort.Sort(Int64Slice(a))
}

func RInt64s(a []int64) {
	sort.Sort(sort.Reverse(Int64Slice(a)))
	// 等价于
	//sort.Slice(Int64Slice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Float64s(a []float64) {
	sort.Sort(Float64Slice(a))
}

func RFloat64s(a []float64) {
	sort.Sort(sort.Reverse(Float64Slice(a)))
	// 等价于
	//sort.Slice(Float64Slice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}

func Strings(a []string) {
	sort.Sort(StringSlice(a))
}

func RStrings(a []string) {
	sort.Sort(sort.Reverse(StringSlice(a)))
	// 等价于
	//sort.Slice(StringSlice(a), func(i, j int) bool {
	//	return a[i] > a[j]
	//})
}
