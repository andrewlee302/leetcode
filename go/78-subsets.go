func subsets(nums []int) [][]int {
    ret := make([][]int, 0, 1<<uint(len(nums)))
    buf := make([]int, len(nums))
    backtracking(nums, 0, buf, 0, &ret)
    return ret
}

func backtracking(nums []int, idx int, buf []int, bufIdx int, ret *[][]int) {
    if len(nums) == idx {
        subset := make([]int, bufIdx)
        copy(subset, buf[:bufIdx])
        *ret = append(*ret, subset)
        return
    }
    buf[bufIdx] = nums[idx]
    backtracking(nums, idx+1, buf, bufIdx+1, ret)
    backtracking(nums, idx+1, buf, bufIdx, ret)
}
