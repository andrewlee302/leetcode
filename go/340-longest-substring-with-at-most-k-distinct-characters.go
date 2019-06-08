// sliding window (longest)
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    left, right := 0, 0 
    dict := make(map[byte]int)
    maxLen := 0
    for right < len(s) {
        dict[s[right]]++
        if len(dict) <= k {
            maxLen = max(maxLen, right-left+1)
        } else {
            for len(dict) > k {
                dict[s[left]]--
                if dict[s[left]] == 0 {
                    delete(dict, s[left]) 
                }
                left++
            }
        }
        right++
    }
    return maxLen
}

func min(a, b int) int { if a < b { return a } else { return b } }
func max(a, b int) int { if a < b { return b } else { return a } }
