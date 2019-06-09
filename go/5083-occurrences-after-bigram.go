import "strings"
func findOcurrences(text string, first string, second string) []string {
    strs := strings.Split(text, " ")
    var ret []string
    for i := 0; i < len(strs); i++ {
        if first == strs[i] && i+1 < len(strs) && second == strs[i+1] && i+2 < len(strs) {
            ret = append(ret, strs[i+2])
        }
    }
    return ret
}
