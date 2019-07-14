func longestWPI(hours []int) int {
    left, right := 0, 0
    diff := 0
    ret := 0
    for left < len(hours) {
        found := false
        maxValid := []int{-1, 0}
        for ; right < len(hours); right++ {
            if hours[right] > 8 { diff++ } else { diff-- }
            if diff > 0 {
                maxValid = []int{right, diff}
                found = true
            }
        }
        if found { 
            right, diff = maxValid[0] + 1, maxValid[1]
            ret = max(ret, right-left) 
            for left < right {
                if hours[left] > 8 { diff-- } else { diff++ }
                left++
                if diff > 0 { break }
            }
        } else { 
            diff = 0
            left++
            right = left
        }
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
