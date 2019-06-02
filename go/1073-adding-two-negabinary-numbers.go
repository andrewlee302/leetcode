func addNegabinary(arr1 []int, arr2 []int) []int {
    maxLen := len(arr1)
    if len(arr2) > len(arr1) { maxLen = len(arr2) }
    if maxLen == 1 && arr1[0] == 0 && arr2[0] == 0 { return []int{0} }
    ret := make([]int, maxLen + 2)
    ptr := len(ret) - 1
    for i1, i2 := len(arr1) - 1, len(arr2) - 1; i1 >= 0 || i2 >= 0; {
        sum := ret[ptr]
        if i1 >= 0 { 
            sum += arr1[i1] 
            i1--
        }
        if i2 >= 0 { 
            sum += arr2[i2] 
            i2--
        }
        if sum == 0 || sum == 1 { 
            ret[ptr] = sum 
        } else if sum == 2 { 
            ret[ptr] = 0 
            ret[ptr-1] += -1
        } else if sum == 3 {
            ret[ptr] = 1 
            ret[ptr-1] += -1
        } else { // sum == -1
            ret[ptr] = 1
            ret[ptr-1] += 1
        }
        fmt.Println(ptr, sum, ret)
        ptr--
    }
    if ret[ptr] == -1 {
        ret[ptr] = 1
        ret[ptr-1] += 1
    }
    i := 0
    for ; i < len(ret); i++ {
        if ret[i] != 0 { break }
    }
    if i == len(ret) { return []int{0} } else { return ret[i:] }
}
