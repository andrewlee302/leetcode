import "sort"
func threeSum(nums []int) [][]int {
    var ret [][]int
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; {
        target := -nums[i]
        for left, right := i + 1, len(nums) - 1; left < right; {
            _2sum := nums[left] + nums[right]
            if _2sum < target {
                for left++; left < right && nums[left-1] == nums[left]; left++ { }
            } else if _2sum > target {
                for right--; left < right && nums[right+1] == nums[right]; right-- { }
            } else {
                ret = append(ret, []int{nums[i],nums[left],nums[right]})
                for left++; left < right && nums[left-1] == nums[left]; left++ { }
                for right--; left < right && nums[right+1] == nums[right]; right-- { }
            }
        }
        for i++; i < len(nums) && nums[i-1] == nums[i]; i++ { }
    }
    return ret
}
