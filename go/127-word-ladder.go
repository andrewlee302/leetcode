import "container/list"
func ladderLength(beginWord string, endWord string, wordList []string) int {
    L := len(beginWord)
    wordsDict := make(map[string][]string)
    // preprocess
    strBuff := make([]byte, L)
    for _, word := range wordList {
        copy(strBuff, word)
        for i := 0; i < L; i++ {
            ws := wordsDict[matchWord(strBuff, word, i)]
            ws = append(ws, word)
            wordsDict[matchWord(strBuff, word, i)] = ws
        }
    }
    
    // find the path
    visited := make(map[string]bool)
    queue := list.New()
    queue.PushBack(beginWord)
    level := 1
    for queue.Len() != 0 {
        size := queue.Len()
        for k := 0; k < size; k++ {
            e := queue.Front()
            word := e.Value.(string)
            queue.Remove(e)
            copy(strBuff, word)
            for i:= 0; i < L; i++ {
                ws, ok := wordsDict[matchWord(strBuff, word, i)]
                if ok {
                    for _, w := range ws {
                        if w == endWord {
                            return level+1
                        }
                        if _, ok := visited[w]; !ok {
                            queue.PushBack(w)  
                            visited[w] = true
                        }
                    }
                }
            }
        }
        level++     
    }
    return 0
}

// wordBytes is the bytes of word, so recovery after constructing the string
func matchWord(wordBytes []byte, word string, idx int) string {
    wordBytes[idx] = '*'
    ret := string(wordBytes)
    wordBytes[idx] = word[idx]
    return ret
}
