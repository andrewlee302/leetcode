func isStrobogrammatic(num string) bool {
    for i, j := 0, len(num) - 1; i < j; {
        if !isUpsideDown(num[i], num[j]) {
            return false
        }
        i++
        j--
    }
    if len(num) % 2 == 1 {
        return isUpsideDown(num[len(num)/2], num[len(num)/2])
    }
    return true
}

func isUpsideDown(a, b byte) bool {
    switch (a) {
        case '0', '1', '8': return b == a
        case '6': return b == '9'
        case '9': return b == '6'
        default: return false
    }
}
