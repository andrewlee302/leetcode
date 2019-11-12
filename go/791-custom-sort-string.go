func customSortString(S string, T string) string {
    var cnts [26]int
    for i := 0; i < len(T); i++ {
        cnts[T[i]-'a']++
    }
    retBs := make([]byte, 0, len(T))
    for i := 0; i < len(S); i++ {
        for n := cnts[S[i]-'a']; n > 0; n-- {
            retBs = append(retBs, S[i])
        }
        cnts[S[i]-'a'] = 0
    }
    for i := 0; i < len(cnts); i++ {
        for n := cnts[i]; n > 0; n-- {
            retBs = append(retBs, byte(i) + 'a')
        }
    }
    return string(retBs)
}
