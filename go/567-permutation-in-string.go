func checkInclusion(s1 string, s2 string) bool {
    dict := make(map[byte]int)
    for i := 0; i < len(s1); i++ { dict[s1[i]]++ }
    substrMatchCnt := 0
    left, right := 0, 0
    for right < len(s2) {
        if occurs, ok := dict[s2[right]]; ok {
            dict[s2[right]] = occurs - 1
            if occurs - 1 == 0 { substrMatchCnt++ }
        }    
        for substrMatchCnt == len(dict) {
            // left mustn't be larger than right.
            if occurs, ok := dict[s2[left]]; ok {
                dict[s2[left]] = occurs + 1
                if occurs == 0 { 
                    substrMatchCnt--
                    if right - left + 1 == len(s1) {
                        return true
                    }
                }
            }
            left++
        }
        right++
    }
    return false
}
