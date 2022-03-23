package pkg

func removeRepeatElement(list []int) []int {
	// 创建一个临时map用来存储数组元素
	temp := make(map[int]struct{})
	for _, v := range list {
		temp[v] = struct{}{}
	}
	var result []int
	for s := range temp {
		result = append(result, s)
	}
	return result
}
