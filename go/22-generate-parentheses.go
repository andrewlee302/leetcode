// Counter diff.
func generateParenthesis(n int) []string {
    buf := make([]byte, n*2)
    var ret []string
    recursion(n, n, 0, buf, 0, &ret)
    return ret
}

// must have diff >= 0
func recursion(leftRemain, rightRemain, diff int, buf []byte, bufIdx int, ret *[]string) {
    if bufIdx == len(buf) {
        *ret = append(*ret, string(buf))
        return
    }
    if rightRemain > 0 && diff > 0 {
        buf[bufIdx] = ')'
        recursion(leftRemain, rightRemain-1, diff-1, buf, bufIdx+1, ret)
    }
    if leftRemain > 0 {
        buf[bufIdx] = '('
        recursion(leftRemain-1, rightRemain, diff+1, buf, bufIdx+1, ret)
    }
}
