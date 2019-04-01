import "strings"
func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 { return nil }
    m := make(map[int][]string)
    for i := 0; i < len(strs); i++ {
        hash := strHash(strs[i])
        m[hash] = append(m[hash], strs[i])
    }
    var ret [][]string
    for _, ss := range m {
        ret = append(ret, ss)
    }
    return ret
}

func strHash(s string) int{
    prime := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
    hash := 1
    for j := 0; j < len(s); j++ { hash *= prime[int(s[j] - 'a')]}
    return hash
}
