func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    currSum, maxSum := nums[0], nums[0]
    
    for i, num := range nums {
        if i == 0 {
            continue
        }
        
        if currSum += num; maxSum < 0 {
            if num > maxSum {
                currSum = num
            }
        } else {
            if currSum < 0 {
                currSum = 0
                continue
            }
        }
        
        maxSum = max(maxSum, currSum)
    }
    
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}