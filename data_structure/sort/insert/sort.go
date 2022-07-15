package insert

func Sort(input []int) []int {
	for i := 1; i < len(input); i++ {
		insertVal := input[i]
		insertIndex := i - 1
		for insertIndex >= 0 && insertVal < input[insertIndex] {
			input[insertIndex+1] = input[insertIndex]
			insertIndex--
		}
		input[insertIndex+1] = insertVal
	}
	return input
}
