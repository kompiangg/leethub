func twoSum(nums []int, target int) []int {
    reverseMap := map[int]int{}
    for index, num := range nums {
        if value, ok := reverseMap[target-num]; ok {
            return []int{value, index}
        }
        reverseMap[num] = index
    }
    return nil
}