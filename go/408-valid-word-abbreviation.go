func validWordAbbreviation(word string, abbr string) bool {
    pos := 0
    for i := 0; i < len(abbr); {
        if abbr[i] >= '1' && abbr[i] <= '9' {
            v := 0
            for ; i < len(abbr) && abbr[i] >= '0' && abbr[i] <= '9'; i++ {
                v = v*10 + int(abbr[i] - '0')
            }
            pos += v
        } else if abbr[i] >= 'a' && abbr[i] <= 'z' {
            if pos >= len(word) || abbr[i] != word[pos] {
                return false
            }
            i++
            pos++
        } else { // abbr[i] == '0'
            return false
        }
    }
    return pos == len(word)
}
