// simple template.
func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(t) > len(s) { return "" }
    left, right := 0, 0
    dictT := make(map[byte]int)
    for i := 0; i < len(t); i++ { dictT[t[i]]++ }
    sizeT := len(dictT)
    ret, retL := "", 0
    for right < len(s) {
        if n, ok := dictT[s[right]]; ok {
            dictT[s[right]]--
            if n == 1 { sizeT-- }
        }
        for sizeT == 0 {
            if n, ok := dictT[s[left]]; ok {
                dictT[s[left]]++
                if n == 0 { sizeT++ }
            }
            if l := right - left + 1; retL == 0 || l < retL {
                ret, retL = s[left: right+1], l
            }
            left++
        }
        right++
    }
    return ret
}

// more aggressive
func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(t) > len(s) { return "" }
    left, right := 0, 0
    dictT := make(map[byte]int)
    for i := 0; i < len(t); i++ { dictT[t[i]]++ }
    sizeT := len(dictT)
    ret, retL := "", 0
    for right < len(s) {
        if n, ok := dictT[s[right]]; ok {
            dictT[s[right]]--
            if n == 1 { sizeT-- }
        }
        for sizeT == 0 {
            if n, ok := dictT[s[left]]; ok {
                dictT[s[left]]++
                if n == 0 {
                    sizeT++
										// aggressive here.
                    if l := right - left + 1; retL == 0 || l < retL {
                        ret, retL = s[left: right+1], l
                    }
                }
            }
            left++
        }
        right++
    }
    return ret
}
