func maxProfit(prices []int) int {
    currMax, maxMax := 0, 0
    
    for i := 1 ; i < len(prices); i++ {
        currMax = max(0, currMax + prices[i] - prices[i-1])
        maxMax = max(maxMax, currMax)
    }
    
    return maxMax
}

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}