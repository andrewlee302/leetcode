func sortColors(nums []int)  {
    i, j, k := 0, 0, len(nums) - 1
    for ; j <= k;  { // wrong: j < k, if [2,0,1]
        if nums[j] == 0 { 
            nums[i], nums[j] = nums[j], nums[i] 
            i++
            j++
        } else if nums[j] == 2 {
            nums[k], nums[j] = nums[j], nums[k]
            k--
        } else {
            j++
        }
    }
}
