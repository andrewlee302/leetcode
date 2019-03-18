func lengthOfLongestSubstringTwoDistinct(s string) int {
    if len(s) < 3 {return len(s) }
    dict := make(map[byte]int)
    cnt := 0
    left, right := 0, 0
    longestLen := 2
    for right < len(s) {
        dict[s[right]]++
        if dict[s[right]] == 1 { cnt++ } // new char
        for cnt == 3 {
            dict[s[left]]--
            if dict[s[left]] == 0 { 
                cnt--
            }
            left++
        }
        if right - left + 1 > longestLen {
            longestLen = right - left + 1
        }
        right++ 
    }
    return longestLen
}
