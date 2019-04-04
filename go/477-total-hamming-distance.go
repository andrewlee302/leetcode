func totalHammingDistance(nums []int) int {
    sum := 0
    for {
        zeroNum := 0
        zeroBitCnt := 0
        for i := 0; i < len(nums); i++ {
            if nums[i] == 0 { 
                zeroBitCnt++ 
                zeroNum++
                continue
            }
            if nums[i] & 0x1 == 0 { zeroBitCnt++ }
            nums[i] = nums[i] >> 1
        }
        sum += zeroBitCnt * (len(nums) - zeroBitCnt)
        if zeroNum == len(nums) { return sum }
    }
}
