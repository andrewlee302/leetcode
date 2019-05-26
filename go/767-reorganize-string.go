type Count struct{
    ch byte
    occur int
}

func reorganizeString(S string) string {
    // count frequency
    dict := make(map[byte]int)
    for i := 0; i < len(S); i++ {
        dict[S[i]]++
    }
    var freqs []Count
    for ch, occur := range dict {
        if occur > (len(S)+1)/2 { return "" }
        freqs = append(freqs, Count{ch, occur})
    }
    // in descending order
    sort.Slice(freqs, func (i, j int) bool { return freqs[i].occur > freqs[j].occur })
    ret := make([]byte, len(S))
    slot := 0
    // fill the same numbers into every other slot.
    for _, freq := range freqs {
        for i := 0; i < freq.occur; i++ {
            // if exceed the length, go back to the slot with index 1.
            if slot >= len(S) { slot = 1 }
            ret[slot] = freq.ch
            slot += 2
        }
    }
    return string(ret)
}
