func permute(nums []int) [][]int {
    var ret [][]int
    idxFlag := make([]bool, len(nums)) // Default: false.
    idxBuff := make([]int, len(nums))
    helper(nums, idxBuff, idxFlag, 0, &ret)
    return ret
}

func helper(nums, idxBuff []int, idxFlag []bool, k int, ret *[][]int) {
    if k == len(nums) {
        arr := make([]int, len(nums))
        for i, idx := range idxBuff { arr[i] = nums[idx] }
        *ret = append(*ret, arr)
    }
    for idx, flag := range idxFlag {
        if !flag { 
            idxFlag[idx] = true
            idxBuff[k] = idx
            helper(nums, idxBuff, idxFlag, k+1, ret)        
            idxFlag[idx] = false
        }
    }
}
