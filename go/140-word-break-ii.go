func wordBreak(s string, wordDict []string) []string {
    dict := make(map[string]bool)
    for _, w := range wordDict {
        dict[w] = true
    }
    cache := make(map[int][]string)
    cache[len(s)] = []string{""}
    return recursion(dict, cache, s, 0)
}

func recursion(dict map[string]bool, cache map[int][]string, s string, idx int) []string {
    if v, ok := cache[idx]; ok { return v }
    var ret []string
    for end := idx; end < len(s); end++ {
        if dict[s[idx:end+1]] {
            appendStrs := recursion(dict, cache, s, end+1)
            for _, str := range appendStrs {
                if str == "" {
                    ret = append(ret, s[idx:end+1])
                } else {
                    ret = append(ret, s[idx:end+1]+" "+str)
                }
            }
        }
    }
    cache[idx] = ret
    return ret
}
