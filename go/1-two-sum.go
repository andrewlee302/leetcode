func twoSum(nums []int, target int) []int {
    cnt := 0
    check := false
    var half = 0
    var indices [2]int 
    if target%2 == 0 {
        check = true
        half = target / 2
    }
    
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        m[nums[i]] = i
        if check && nums[i] == half {
            if cnt < 2 {
                indices[cnt] = i
            } else {
                break
            }
            cnt ++
        }
    }
    if cnt == 2 {
        return indices[:]
    }
    
    for i := 0; i < len(nums); i++ {
        if half != nums[i] {
            value, ok := m[target - nums[i]] 
            if ok {
                return []int {i, value}
            }
        }
    }
    return []int{0, 0}
}

