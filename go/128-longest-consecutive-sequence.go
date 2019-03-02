import "sort"
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Ints(nums)
    longest := 1
    currentStreak := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            if nums[i] == nums[i-1] + 1 {
                currentStreak ++
            } else {
                if currentStreak > longest {
                    longest = currentStreak   
                }
                currentStreak = 1
            }
        }
    }
    if currentStreak > longest {
        longest = currentStreak
    }
    return longest
}
