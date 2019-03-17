func subsets(nums []int) [][]int {
    ret := make([][]int, 0, 1<<uint(len(nums)))
    subset := make([]int, len(nums))
    backtracking(subset, 0, nums, &ret)
    return ret
}

func backtracking(subset []int, idx int, nums []int, ret *[][]int) {
    if len(nums) == 0 {
        finalSubset := make([]int, idx)
        copy(finalSubset, subset[:idx])
        *ret = append(*ret, finalSubset)
        return
    }
    subset[idx] = nums[0]
    backtracking(subset, idx+1, nums[1:], ret)
    backtracking(subset, idx, nums[1:], ret)
}
