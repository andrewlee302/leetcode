func addBinary(a string, b string) string {
    ret := make([]byte, max(len(a), len(b)) + 1)
    var carry byte = 0
    for i := 0; i < len(ret); i++ {
        var ai, bi byte = 0, 0
        if i < len(a) { ai = a[len(a)-i-1] - '0' }
        if i < len(b) { bi = b[len(b)-i-1] - '0' }
        sum := ai + bi + carry 
        curr := sum & 1
        carry = (sum - curr) >> 1
        ret[len(ret)-i-1] = curr + '0'
    }
    if ret[0] == '0' { return string(ret[1:]) } else { return string(ret) }
}

func max(i, j int) int { if i > j { return i } else { return j } }
