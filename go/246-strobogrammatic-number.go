func isStrobogrammatic(num string) bool {
    m := map[byte]byte{'6':'9', '9':'6', '1':'1', '8':'8', '0':'0'}
    for l, r := 0, len(num)-1; l <= r; l, r = l+1, r-1 {
        if m[num[l]] != num[r] { return false }
    }
    return true
}
