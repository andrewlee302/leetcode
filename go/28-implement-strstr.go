func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    for start := 0; start <= len(haystack) - len(needle); start ++{
        i, j := start, 0
        for ;j < len(needle) && haystack[i] == needle[j]; {
            i++
            j++
        }
        if j == len(needle) {
            return start
        }
    }
    return -1
}
