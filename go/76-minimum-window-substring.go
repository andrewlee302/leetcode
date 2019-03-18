func minWindow(s string, t string) string {
    if len(t) == 0  || len(s) < len(t) {return ""}
    tDict := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tDict[t[i]]++
    }
    cnt := len(tDict)    
    left, right := 0, 0
    minimalTuples := []int{int(^uint(0)>>1), 0, 0} // length, left, right
    for right < len(s) {
        if remaining, ok := tDict[s[right]]; ok {
            tDict[s[right]] = remaining - 1   
            if tDict[s[right]] == 0 { cnt-- }
        }
        for cnt == 0 {
            // left mustn't be more than right
            if remaining, ok := tDict[s[left]]; ok {
                tDict[s[left]] = remaining + 1   
                if tDict[s[left]] > 0 { 
                    cnt++
                    matchLen := right -left + 1
                    if matchLen < minimalTuples[0] {
                        minimalTuples = []int{matchLen, left, right}
                    }
                }
            }
            left++
        }
        right++
    }
    if minimalTuples[0] != int(^uint(0)>>1) {
        return s[minimalTuples[1]:minimalTuples[2]+1]
    } else {
        return ""
    }
}


