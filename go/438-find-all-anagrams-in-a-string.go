func findAnagrams(s string, t string) []int  {
    if len(t) == 0  || len(s) < len(t) {return []int{}}
    tDict := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tDict[t[i]]++
    }
    cnt := len(tDict)    
    left, right := 0, 0
    var ret []int
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
                    if right - left + 1 == len(t) {
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
