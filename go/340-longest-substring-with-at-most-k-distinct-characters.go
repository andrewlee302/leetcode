// sliding window (longest)
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    left, right := 0, 0
    dict := make(map[byte]int)
    maxLen := 0
    for right < len(s) {
        dict[s[right]]++
        for len(dict) > k {
            dict[s[left]]--
            if dict[s[left]] == 0 {
                delete(dict, s[left])
            }
            left++
        }
        maxLen = max(maxLen, right-left+1)
        right++
    }
    return maxLen
}

func max(a, b int) int { if a < b { return b } else { return a } }

// no map delete
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    if len(s) == 0 || k <= 0 { return 0 }
    left, right := 0, 0
    dict := make(map[byte]int)
    dictSize := 0
    ret := 0
    for right < len(s) {
        dict[s[right]]++
        if dict[s[right]] == 1 { dictSize++ }
        for dictSize > k {
            dict[s[left]]--
            if dict[s[left]] == 0 {
								dictSize--
            }
            left++
        }
        ret = max(ret, right-left+1)
        right++
    }
    return ret
}

func max(i, j int) int { if i > j { return i } else { return j } }
