package removeelement

func removeElement(nums []int, val int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[c] = nums[i]
			c++
		}
	}
	return c
}
