func isPalindrome(s string) bool {
    for left, right := 0, len(s) - 1; left < right;  {
        if !isAlphaOrNumber(s[left]) { left++; continue }
        if !isAlphaOrNumber(s[right]) { right--; continue }
        if isAlpha(s[left]) && isAlpha(s[right]) {
            var diff byte
            if s[left] > s[right] { diff = s[left] - s[right] } else { diff = s[right] - s[left] }
            if diff != 0 && diff != 'a' - 'A' { return false }
        } else if isNumber(s[left]) && isNumber(s[right]) {
            if s[left] != s[right] { return false }
        } else {
            return false
        }
        left, right = left+1, right-1
    }
    return true
}

func isAlphaOrNumber(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' }
func isAlpha(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' }
func isNumber(b byte) bool { return b >= '0' && b <= '9' }
