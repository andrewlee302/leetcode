// length of the final result (non-positive two numbers)
// multiplication: len(num1)+len(num2)-1 or len(num1)+len(num2). (Note: zero is exception, so return directy.)
// sum: max(len(num1), len(num2)) or max(len(num1), len(num2)) + 1 (Note: no exception.)
func addBinary(a string, b string) string {
    carry := 0
    maxLen := len(a)+1
    if len(a) < len(b) { maxLen = len(b)+1 }
    ret := make([]byte, maxLen)
    for i := 0; i < maxLen; i++ {
        aVal, bVal := 0, 0
        if i < len(a) { aVal = int(a[len(a)-i-1] - '0') }
        if i < len(b) { bVal = int(b[len(b)-i-1] - '0') }
        sum := aVal + bVal + carry
        ret[len(ret)-i-1] = byte(sum%2) + '0'
        carry = sum/2
    }
    if ret[0] == '0' { return string(ret[1:]) } else { return string(ret) }
}
