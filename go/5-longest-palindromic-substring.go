// DP.
func longestPalindrome(s string) string {
    if len(s) == 0 { return "" }
    dp := make([][]bool, len(s))
    var ret string
    for i := 0; i < len(s); i++ { dp[i] = make([]bool, len(s)) }
    for l := 1; l <= len(s); l++ {
        for i := 0; i+l-1 < len(s); i++ {
            j := i+l-1
            if l == 1 || l == 2 && s[i] == s[j] || l > 2 && dp[i+1][j-1] && s[i] == s[j] {
                dp[i][j] = true
                ret = s[i:j+1]
            }
        }
    }
    return ret
}

// dp[i][j] true or false
// dp[i][j] =
// if dp[i+1][j-1] == true && s[i] == s[j]: true
// else: false
// dp[i][i] = true
// dp[i][i+1] = if s[i] == s[i+1]: true else: false

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
