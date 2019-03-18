func isPalindrome(s string) bool {
    var diff uint8 = 'a' - 'A'
    i, j := 0, len(s) - 1
    for i < j {
        for ; i < len(s) && !isAlpha(s[i]) && !isNumberic(s[i]); i++ {}
        for ; j >= 0 && !isAlpha(s[j]) && !isNumberic(s[j]); j-- {}
        if i < len(s) && j >= 0 && s[i] != s[j] {
            if isAlpha(s[i]) && s[i] - s[j] != diff && s[i] - s[j] != -diff {   
                return false
            }
            if isNumberic(s[i]) {
                return false
            }
        }
        i++
        j--
    }
    return true
}
    
func isAlpha(b byte) bool {
     return b <= 'z' && b >= 'a' || b <= 'Z' && b >= 'A'   
}
    
func isNumberic(b byte) bool {
    return b <= '9' && b >= '0'
}
