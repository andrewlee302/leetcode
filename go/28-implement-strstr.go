func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    for i := 0; i <= len(haystack) - len(needle); i++ {
        for j := 0; j < len(needle); j++ {
            if haystack[i+j] != needle[j] {
                break
            }
            if j == len(needle) - 1 {
                return i
            }
        }
    }
    return -1
}
