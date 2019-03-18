func isAlienSorted(words []string, order string) bool {
    if len(words) == 0 || len(words) == 1 {
        return true
    }
    m := make([]int, 26)
    for i := 0; i < len(order); i++ {
        m[order[i]-'a'] = i
    }
    for i := 1; i < len(words); i ++ {
        prevLen, currLen := len(words[i-1]), len(words[i])
        minLen := prevLen
        if currLen < minLen {
            minLen = currLen
        }
        j := 0
        for ; j < minLen; j++ {
            prev, curr := m[words[i-1][j]-'a'], m[words[i][j]-'a']
            if prev < curr {
                break
            } else if prev > curr {
                return false
            }
        }
        if j == minLen && prevLen > currLen {
            return false
        }
    }
    return true
}
