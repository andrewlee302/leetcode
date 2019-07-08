import "container/list"
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    wordPatterns := make(map[string][]string)
    for _, word := range wordList {
        wordBytes := []byte(word)
        for i := 0; i < len(word); i++ {
            pattern := getPattern(wordBytes, i)
            wordPatterns[pattern] = append(wordPatterns[pattern], word)
        }
    }
    parentsMap := make(map[string][]string)
    levels := make(map[string]int) // levels
    queue := list.New()
    queue.PushBack(beginWord)
    levels[beginWord] = 1
    
    found := false
    level := 1
    for !found && queue.Len() != 0 {
        size := queue.Len()
        for s := 0; s < size; s++ {
            e := queue.Front()
            val := e.Value.(string)
            queue.Remove(e)
            wordBytes := []byte(val)
            for i := 0; i < len(wordBytes); i++ {
                pattern := getPattern(wordBytes, i)
                if matchWords, ok := wordPatterns[pattern]; ok {
                    for _, w := range matchWords {
                        if w == endWord { found = true }
                        if levels[w] == 0 {
                            levels[w] = level+1
                            queue.PushBack(w)
                        }
                        if levels[w] == level+1 { // NOTE!
                            parentsMap[w] = append(parentsMap[w], val)
                        } 
                    }
                }
            }
        }
        level++
    }
    buf := make([]string, level)
    var ret [][]string
    buf[level-1] = endWord
    fmt.Println(parentsMap)
    helper(buf, level-2, parentsMap, &ret)
    return ret
}

func helper(buf []string, idx int, parentsMap map[string][]string, ret *[][]string) {
    if idx == -1 {
        c := make([]string, len(buf))
        copy(c, buf)
        *ret = append(*ret, c)
        return
    }
    for _, parent := range parentsMap[buf[idx+1]] {
        buf[idx] = parent
        helper(buf, idx-1, parentsMap, ret)
    }
}

func getPattern(word []byte, idx int) string {
    b := word[idx]
    word[idx] = '*'
    ret := string(word)
    word[idx] = b
    return ret
}
