func minAddToMakeValid(S string) int {
    diff := 0
    addCnt := 0
    for i := 0; i < len(S); i++ {
        if S[i] == '(' { diff++ } else { diff-- }
        if diff < 0 {
            addCnt++
            diff++
        }
    }
    return addCnt + diff
}
