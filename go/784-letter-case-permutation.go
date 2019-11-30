// Iteration.
func letterCasePermutation(S string) []string {
    buffers := [][]byte{make([]byte, len(S))}
    for i := 0; i < len(S); i++ {
        v := S[i]
        size := len(buffers)
        for k := 0; k < size; k++ {
            buffers[k][i] = v
            if !(v >= '0' && v <= '9') {
                buf := make([]byte, len(S))
                copy(buf, buffers[k])
                if v >= 'a' && v <= 'z' {
                    buf[i] = v-'a'+'A'
                } else {
                    buf[i] = v-'A'+'a'
                }
                buffers = append(buffers, buf)
            }
        }
    }
    ret := make([]string, len(buffers))
    for i, buf := range buffers { ret[i] = string(buf) }
    return ret
}

// Recursion.
func letterCasePermutation(S string) []string {
    if len(S) == 0 { return []string{""} }
    parts := letterCasePermutation(S[1:])
    var ret []string
    for _, part := range parts {
        ret = append(ret, S[0:1] + part)
        if v := S[0]; v >= 'a' && v <= 'z' {
            ret = append(ret, string([]byte{v-'a'+'A'}) + part)
        } else if v >= 'A' && v <= 'Z' {
            ret = append(ret, string([]byte{v-'A'+'a'}) + part)
        }
    }
    return ret
}
