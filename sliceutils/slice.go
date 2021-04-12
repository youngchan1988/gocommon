package sliceutils

func SumInt(source []int) (sum int) {
	for _, v := range source {
		sum += v
	}
	return
}

func SumInt64(source []int64) (sum int64) {
	for _, v := range source {
		sum += v
	}
	return
}

//int 切片 元素下标
func IndexInt(slice []int, elem int) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//int64 切片 元素下标
func IndexInt64(slice []int64, elem int64) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//string 切片 元素下标
func IndexString(slice []string, elem string) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//int 切片 是否包含元素
func ContainInt(slice []int, elem int) bool {
	index := IndexInt(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//int64 切片 是否包含元素
func ContainInt64(slice []int64, elem int64) bool {
	index := IndexInt64(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//string 切片 是否包含元素
func ContainString(slice []string, elem string) bool {
	index := IndexString(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//int 切片 删除元素
func RemoveInt(slice []int, elem ...int) []int {
	maps := map[int]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []int
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//int64 切片 删除元素
func RemoveInt64(slice []int64, elem ...int64) []int64 {
	maps := map[int64]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []int64
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//string 切片 删除元素
func RemoveString(slice []string, elem ...string) []string {
	maps := map[string]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []string
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//int 切片 元素去重复
func DistinctInt(slice []int) []int {
	maps := map[int]bool{}
	var newSlice []int
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//int64 切片 元素去重复
func DistinctInt64(slice []int64) []int64 {
	maps := map[int64]bool{}
	var newSlice []int64
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//string 切片 元素去重复
func DistinctString(slice []string) []string {
	maps := map[string]bool{}
	var newSlice []string
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}
