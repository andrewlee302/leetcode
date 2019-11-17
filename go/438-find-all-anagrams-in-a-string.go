// Incremental.
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) { return nil }
    dictP := make(map[byte]int)
    for i := 0; i < len(p); i++ {
        dictP[p[i]-'a']++
    }
    matchRemain := len(dictP)
    for i := 0; i < len(p); i++ {
        if v, ok := dictP[s[i]-'a']; ok {
            dictP[s[i]-'a'] = v-1
            if v - 1 == 0 { matchRemain-- }
        }
    }
    var ret []int
    if matchRemain == 0 { ret = append(ret, 0) }
    for i := 1; i <= len(s) - len(p); i++ {
        if v, ok := dictP[s[i+len(p)-1]-'a']; ok {
            dictP[s[i+len(p)-1]-'a'] = v-1
            if v - 1 == 0 { matchRemain-- }
        }
        if v, ok := dictP[s[i-1]-'a']; ok {
            dictP[s[i-1]-'a'] = v+1
            if v + 1 == 1 { matchRemain++ }
        }
        if matchRemain == 0 { ret = append(ret, i) }
    }
    return ret
}

// Window template.
func findAnagrams(s string, p string) []int {
    if len(s) < len(p) { return nil }
    pDict := make(map[byte]int)
    for i := 0; i < len(p); i++ { pDict[p[i]]++ }
    subStringMatchCnt := 0
    left, right := 0, 0
    var ret []int
    for right < len(s) {
        if remaining, ok := pDict[s[right]]; ok {
            pDict[s[right]] = remaining - 1
            if remaining - 1 == 0 { subStringMatchCnt++ }
        }
        for subStringMatchCnt == len(pDict) {
            // left mustn't be more than right
            if remaining, ok := pDict[s[left]]; ok {
                pDict[s[left]] = remaining + 1
                if remaining == 0 {
                    subStringMatchCnt--
                    // left~right
                    if right - left + 1 == len(p) {
                        ret = append(ret, left)
                    }
                }
            }
            left++
        }
        right++
    }
    return ret
}
