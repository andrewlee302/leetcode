func minRemoveToMakeValid(s string) string {
    sb := []byte(s)
    left := 0
    for i := 0; i < len(sb); i++ {
        if sb[i] == '('  {
            left++
        } else if sb[i] == ')' {
            if left > 0 { left-- } else { sb[i] = 0 }
        }
    }
    right := 0
    for i := len(sb) - 1; i >= 0; i-- {
        if sb[i] == ')'  {
            right++
        } else if sb[i] == '(' {
            if right > 0 { right-- } else { sb[i] = 0 }
        }
    }
    retB := make([]byte, 0, len(s))
    for _, b := range sb { if b != 0 { retB = append(retB, b) } }
    return string(retB)
}
