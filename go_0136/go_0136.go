package go0136

func singleNumber(nums []int) int {
	var aa int
	for i := range nums {
		aa = aa ^ nums[i]
	}

	return aa
}
