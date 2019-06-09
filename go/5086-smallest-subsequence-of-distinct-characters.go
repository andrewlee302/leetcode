// greedy + recursion.
import "strings"
func smallestSubsequence(text string) string {
    if text == "" { return "" }
    dict := make(map[byte]int)
    for i := 0; i < len(text); i++ { dict[text[i]]++ }
    minPos := 0
    for i := 0; i < len(text); i++ {
        if text[i] < text[minPos] { minPos = i }
        dict[text[i]]--
        if dict[text[i]] == 0 { break }
    }
    minStr := string([]byte{text[minPos]})
    restStr := strings.Replace(text[minPos+1:], minStr, "", -1)
    return minStr + smallestSubsequence(restStr)
}
