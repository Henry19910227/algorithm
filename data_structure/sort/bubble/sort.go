package bubble

func Sort(input []int) []int {
	for i := 0; i < len(input); i++ {
		// 先減去一表示最後一個元素不需繼續排序
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				tmp := input[j+1]
				input[j+1] = input[j]
				input[j] = tmp
			}
		}
	}
	return input
}
