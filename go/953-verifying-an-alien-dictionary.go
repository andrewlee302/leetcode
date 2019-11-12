func isAlienSorted(words []string, order string) bool {
    if len(words) <= 1 { return true }
    var orderInvert [26]int
    for i := 0; i < len(order); i++ {
        orderInvert[order[i]-'a'] = i
    }
    for i := 1; i < len(words); i++ {
        s1, s2 := words[i-1], words[i]
        l := min(len(s1), len(s2))
        equal := true
        for j := 0; j < l; j++ {
            o1, o2 := orderInvert[s1[j]-'a'], orderInvert[s2[j]-'a']
            if o1 > o2 {
                return false
            } else if o1 < o2 {
                equal = false
                break
            }
        }
        if equal && len(s1) > len(s2) { return false }
    }
    return true
}

func min(i, j int) int { if i < j { return i } else { return j } }
