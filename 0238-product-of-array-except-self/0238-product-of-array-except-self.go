func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}

	res[0] = 1
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	right := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right = right * nums[i]
	}

	return res
}