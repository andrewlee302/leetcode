func findSubstring(s string, words []string) []int {
    if len(s) == 0 || len(words) == 0 || len(s) < len(words[0]) * len(words) { return []int{}}
    unit := len(words[0])
    numWords := len(words)
    dict := make(map[string]int)
    for _, w := range words {
        dict[w]++
    }
    dictSize := len(dict)
    var ret []int
    for i := 0; i < unit; i++ {
        // the inner code block is like leetcode 438.
        currDict := make(map[string]int)
        cnt := 0
        for left, right := i, i; right <= len(s) - unit; right += unit {
            subStr := s[right:right+unit]
            if requiredOccur, ok := dict[subStr]; ok {
                occur := currDict[subStr] + 1
                currDict[subStr] = occur
                if occur == requiredOccur {
                    cnt++
                }
                for cnt == dictSize {
                    subStr = s[left:left+unit]
                    if requiredOccur, ok = dict[subStr]; ok {
                        occur = currDict[subStr] - 1
                        currDict[subStr] = occur
                        if occur < requiredOccur {
                            cnt--
                            if right - left + unit == unit*numWords {
                                ret = append(ret, left)
                            }
                        }
                    }
                    left += unit
                }
            }
        }
    }
    return ret
}
