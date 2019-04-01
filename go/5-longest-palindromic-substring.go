func longestPalindrome(s string) string {
    if len(s) == 0 { return "" }
    maxLen := 1
    ret := []int{0, 0}
    for i := 0; i < len(s) - 1; i++ {
        leftOdd, rightOdd := expand(s, i, i)
        lenOdd := rightOdd - leftOdd + 1
        if maxLen < lenOdd {
            maxLen = lenOdd
            ret = []int{leftOdd, rightOdd}
        }
        leftEven, rightEven := expand(s, i, i + 1)
        lenEven := rightEven - leftEven + 1
        if maxLen < lenEven {
            maxLen = lenEven
            ret = []int{leftEven, rightEven}
        }
    }
    return s[ret[0]: ret[1]+1]
}

func expand(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left - 1, right + 1 { }
    return left + 1, right - 1
}
