package base

// Swap 兩數交換
func Swap(a *int, b *int)  {
	tmp := *a
	*a = *b
	*b = tmp
}

// MoveBack 整體向後位移一位
func MoveBack(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	last := nums[len(nums) - 1]
	index := len(nums) - 1
	for index >= 0 {
		nums[index - 1] = nums[index]
		index --
	}
	nums[0] = last
	return nil
}
