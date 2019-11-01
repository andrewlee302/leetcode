// window.
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 { return 0 }
    left, right := 0, 0
    counter := make(map[byte]int)
    ret := 0
    for right < len(s) {
        counter[s[right]]++
        for counter[s[right]] == 2 {
            ret = max(ret, right-left)
            counter[s[left]]--
            left++
        }
        right++
    }
    ret = max(ret, right-left)
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }

// DP
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 { return 0 }
    pos := make(map[byte]int)
    ret := 0
    lastDistance := 0
    for i := 0; i < len(s); i++ {
        dist := lastDistance + 1
        if idx, ok := pos[s[i]]; ok {
            dist = min(dist, i - idx)
        }
        ret = max(ret, dist)
        lastDistance = dist
        pos[s[i]] = i
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
func min(i, j int) int { if i < j { return i } else { return j } }
