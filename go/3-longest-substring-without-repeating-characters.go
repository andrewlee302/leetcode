func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    dict := make(map[byte]int)
    left, right := 0, 0
    longestLen := 0
    for right < len(s) {
        if _, ok := dict[s[right]]; !ok {
            dict[s[right]] = 1
            if right - left + 1 > longestLen {
                longestLen = right - left + 1
            }
        } else {
            for s[left] != s[right] {
                delete(dict, s[left])
                left++
            }
            left++
        } 
        right++
    }
    return longestLen  
}
