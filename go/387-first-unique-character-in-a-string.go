func firstUniqChar(s string) int {
    m := make([]int, 26) 
    for i := 0; i < len(s); i++ { m[int(s[i]-'a')]++ }
    for i := 0; i < len(s); i++ { if m[int(s[i]-'a')] == 1 { return i } }
    return -1
}
