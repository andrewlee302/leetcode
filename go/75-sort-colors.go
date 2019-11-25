func sortColors(nums []int)  {
    if len(nums) <= 1 { return }
    p1, p2 := 0, 0
    for ; p2 < len(nums); p2++ {
        if nums[p2] == 0 {
            nums[p1], nums[p2] = nums[p2], nums[p1]
            p1++
        }
    }
    p2 = p1
    for ; p2 < len(nums); p2++ {
        if nums[p2] == 1 {
            nums[p1], nums[p2] = nums[p2], nums[p1]
            p1++
        }
    }
    return
}
