// bi-BFS
import "container/list"
func ladderLength(beginWord string, endWord string, wordList []string) int {
    if len(wordList) == 0 { return 0 }
    wildcardMap := make(map[string][]string)
    
    // Note!!!. if endWord is not in list, return 0.
    flag := false
    for _, word := range wordList {
        if word == endWord { flag = true } 
        for _, wild := range allWilds(word) {
            wildcardMap[wild] = append(wildcardMap[wild], word)
        }  
    }
    if !flag { return 0 }
    
    visited1 := make(map[string]bool)
    visited2 := make(map[string]bool)
    queue1, queue2 := list.New(), list.New()
    queue1.PushBack(beginWord)
    visited1[beginWord] = true
    queue2.PushBack(endWord)
    visited2[endWord] = true
    depth1, depth2 := 0, 0
    for queue1.Len() != 0 && queue2.Len() != 0 {
        size := queue1.Len()
        for i := 0; i < size; i++ {
            e := queue1.Front()
            queue1.Remove(e)
            v := e.Value.(string)
            for _, wild := range allWilds(v) {
                for _, nextWord := range wildcardMap[wild] {
                    if !visited1[nextWord] {
                        if !visited2[nextWord] {
                            visited1[nextWord] = true
                            queue1.PushBack(nextWord)
                        } else {
                            return depth1 + 1 + depth2 + 1
                        }
                    }
                }
            }
        }
        depth1 += 1
        
        size = queue2.Len()
        for i := 0; i < size; i++ {
            e := queue2.Front()
            queue2.Remove(e)
            v := e.Value.(string)
            for _, wild := range allWilds(v) {
                for _, nextWord := range wildcardMap[wild] {
                    if !visited2[nextWord] {
                        if !visited1[nextWord] {
                            visited2[nextWord] = true
                            queue2.PushBack(nextWord)
                        } else {
                            return depth1 + 1 + depth2 + 1
                        }
                    }
                }
            }
        }
        depth2 += 1
    }
    return 0
}

func allWilds(word string) []string {
    var ret []string
    wb := []byte(word)
    for i := 0; i < len(word); i++ {
        wb[i] = '*'
        wild := string(wb)
        ret = append(ret, wild)
        wb[i] = word[i]
    }
    return ret
}

// normal BFS
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
