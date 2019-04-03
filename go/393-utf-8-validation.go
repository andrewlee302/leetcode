func validUtf8(data []int) bool {
    m := map[int]int{7:1,5:2,4:3,3:4}
    size := 0
    for i := 0; i < len(data); i += size {
        mask := 1 << 7
        j := 7
        for ; j >= 3; j-- {
            if data[i] & mask == 0 { break }
            mask = mask >> 1
        }
        if m[j] == 0 { return false }
        size = m[j]
        if i + size - 1 >= len(data) { return false } // Note
        for k := 1; k <= size - 1; k++ {
            if data[i+k] & 0xC0 != 0x80 { return false }
        }
    }
    return true
}
