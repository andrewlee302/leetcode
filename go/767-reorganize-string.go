import "sort"
func reorganizeString(S string) string {
    if len(S) == 0 { return "" }
    var dict [26][2]int
    for i := 0; i < len(dict); i++ { dict[i][0] = i }
    for i := 0; i < len(S); i++ { dict[S[i]-'a'][1]++ }
    // in descending order
    sort.Slice(dict[:], func(i, j int) bool { return dict[i][1] > dict[j][1] })
    if dict[0][1] > (len(S)+1)/2 { return "" }
    buf := make([]byte, len(S))
    slot := 0
    // fill the same numbers into every other slot.
    for _, v := range dict {
        b := byte(v[0]) + 'a'
        for i := 0; i < v[1]; i++ {
            // if exceed the length, go back to the slot with index 1.
            if slot >= len(S) { slot = 1 }
            buf[slot] = b
            slot += 2
        }
    }
    return string(buf)
}
