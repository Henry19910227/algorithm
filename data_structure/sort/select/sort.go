package _select

func Sort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		var min = input[i+1]
		var minIndex = i + 1
		for j := i + 1; j < len(input)-1; j++ {
			if input[j] < min {
				min = input[j]
				minIndex = j
			}
		}
		if min < input[i] {
			tmp := input[i]
			input[i] = input[minIndex]
			input[minIndex] = tmp
		}
	}
	return input
}
